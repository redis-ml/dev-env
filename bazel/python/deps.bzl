load("@rules_python//python:pip.bzl", _pip_parse="pip_parse")
load("@python3_10//:defs.bzl", "interpreter")

def pip_parse(name, interpreter=interpreter, **kwargs):
    _pip_parse(
        name=name,
        python_interpreter_target=interpreter,
        **kwargs,
    )


def install_all_pip_deps():
    # Create a central external repo, @my_deps, that contains Bazel targets for all the
    # third-party packages specified in the requirements.txt file.
    pip_parse(
        name = "pip_deps",
        requirements = "//bazel/python:requirements.txt",
    )
    pip_parse(
        name = "demo_bazel_pip_deps",
        requirements = "//python/demo_bazel:requirements_lock.txt",
    )
    pip_parse(
        name = "django_hello_world_pip_deps",
        requirements = "//:python/demo_bazel/django_hello_world/requirements.txt",
    )
    pip_parse(
        name = "pip_jupyter_deps",
        requirements = "//python/jupyter:requirements_lock.txt",
    )
