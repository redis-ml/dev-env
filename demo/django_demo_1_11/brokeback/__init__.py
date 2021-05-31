# isort:skip_file
from __future__ import absolute_import

# This will make sure the app is always imported when
# Django starts so that shared_task will use this app.
from brokeback.celeryapp import app as celery_app

# Patch DecimalField.quantize to circumvent DRF CDecimal incompatibility
# See https://phabricator.robinhood.com/D41442
import utils.fields

utils.fields.patch_decimal_field_quantize()

__all__ = ["celery_app"]
