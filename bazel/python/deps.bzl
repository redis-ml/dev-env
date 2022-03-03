load("@rules_python//python:pip.bzl", "pip_install")


def install_all_pip_deps():
  # Create a central external repo, @my_deps, that contains Bazel targets for all the
  # third-party packages specified in the requirements.txt file.
  pip_install(
     name = "pip_deps",
     requirements = "//bazel/python:requirements.txt",
  )
  pip_install(
     name = "demo_bazel_pip_deps",
     requirements = "//:python/demo_bazel/requirements.txt",
  )
  pip_install(
     name = "django_hello_world_pip_deps",
     requirements = "//:python/demo_bazel/django_hello_world/requirements.txt",
  )

