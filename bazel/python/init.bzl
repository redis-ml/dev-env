load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def bootstrap_python_rules():
  http_archive(
      name = "rules_python",
      sha256 = "a30abdfc7126d497a7698c29c46ea9901c6392d6ed315171a6df5ce433aa4502",
      strip_prefix = "rules_python-0.6.0",
      url = "https://github.com/bazelbuild/rules_python/archive/0.6.0.tar.gz",
  )

  ########################
  # Example of an "unreleased version"
  # For details, check https://github.com/bazelbuild/rules_python .

  # rules_python_version = "740825b7f74930c62f44af95c9a4c1bd428d2c53" # Latest @ 2021-06-23

  # http_archive(
  #     name = "rules_python",
  #     sha256 = "3474c5815da4cb003ff22811a36a11894927eda1c2e64bf2dac63e914bfdf30f",
  #     strip_prefix = "rules_python-{}".format(rules_python_version),
  #     url = "https://github.com/bazelbuild/rules_python/archive/{}.zip".format(rules_python_version),
  # )
