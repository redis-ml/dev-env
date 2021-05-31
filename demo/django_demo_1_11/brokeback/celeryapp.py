# isort:skip_file
from __future__ import absolute_import

from utils.backcompat import override_decimal

override_decimal()  # noqa

import arrow
from contextlib import contextmanager

import celery
from celery.utils import cached_property
from celery.utils.imports import instantiate
from celery.five import monotonic

from django.conf import settings
from django.db import transaction

from utils.statsd import stats_client
from zardoz.utils.log import get_logger

# Hack: to import tracing, we need to run the path setup that happens in
# settings.py. Accessing a property on settings forces that module to be
# imported
settings.USE_CELERY
from tracing.tracer import tracer_for  # noqa

LOGGER = get_logger(__name__)

TASK_CREATED_AT_HEADER = "_task_created_at"


class CeleryCallInTransaction(Exception):
    def __init__(self):
        super(CeleryCallInTransaction, self).__init__(
            "Calling celery task inside of django transaction"
        )


class BBTask(celery.Task):
    abstract = True

    def apply_async(self, *args, **kwargs):
        kwargs["headers"] = kwargs.get("headers", {})
        kwargs["headers"][TASK_CREATED_AT_HEADER] = str(arrow.now())

        with tracer_for("celery").start_span("publish") as span:
            span.set_tag("name", self.name)
            # Log times where we spin off a celery task inside of transaction
            # -- we  want to use transaction.on_commit instead.
            # Use an exception so the full stack trace gets logged to sentry.
            try:
                if transaction.get_connection().in_atomic_block:
                    raise CeleryCallInTransaction()
            except CeleryCallInTransaction as e:
                LOGGER.exception(e)
            self._check_args(*args, **kwargs)
            self.stats_client.incr(self.name)
            with self.log_time_elapsed("celery-task-publish"):
                return super(BBTask, self).apply_async(*args, **kwargs)

    def __call__(self, *args, **kwargs):
        """ __call__() get executed by the Brokeback Celery workers whenever
        a task gets executed."""
        self.stats_client.incr("worker_{}".format(self.name))
        created_at = self.get_header(TASK_CREATED_AT_HEADER)
        if created_at:
            elapsed = (arrow.now() - arrow.get(created_at)).total_seconds()
            self.stats_client.timing(
                "worker_received_latency_{}".format(self.name), elapsed * 1000
            )
        return super(BBTask, self).__call__(*args, **kwargs)

    def _check_args(self, *args, **kwargs):
        if settings.CELERY_ARG_VALIDATE_FUNC:
            settings.CELERY_ARG_VALIDATE_FUNC(*args, **kwargs)

    def get_header(self, key, default=None):
        return (self.request.headers or {}).get(key, default)

    @contextmanager
    def log_time_elapsed(self, statsd_name):
        start_time = monotonic()
        yield
        self.stats_client.timing(statsd_name, monotonic() - start_time)

    @cached_property
    def stats_client(self):
        return stats_client("celery")


class App(celery.Celery):
    @cached_property
    def amqp(self):
        amqp_cls = getattr(settings, "CELERY_AMQP_CLS", None)
        if amqp_cls is not None:
            return instantiate(amqp_cls, app=self)
        return instantiate(self.amqp_cls, app=self)


app = App("brokeback", task_cls=BBTask,)


# Using a string here means the worker will not have to
# pickle the object when using Windows.
app.config_from_object("django.conf:settings")
app.autodiscover_tasks(lambda: settings.INSTALLED_APPS, related_name="tasks")

# also autodiscover tasks from notifications.py
app.autodiscover_tasks(
    lambda: settings.INSTALLED_APPS, related_name="notifications"
)
