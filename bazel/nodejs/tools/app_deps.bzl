
# The npm_install rule runs yarn anytime the package.json or package-lock.json file changes.
# It also extracts any Bazel rules distributed in an npm package.
load("@build_bazel_rules_nodejs//:index.bzl", "npm_install", "yarn_install")

_ALL_APP_PACKAGES = {
  "npm": [
      "yarn",
      # These files MUST be placed in the Workspace Root.
      "//:package.json",
      "//:yarn.lock",
  ],
  # Used by "demo_bazel" project.
  "npm_demo_bazel": [
      "yarn",
      "//js/demo_bazel:package.json",
      "//js/demo_bazel:yarn.lock",
  ],
}

def install_nodejs_app_dependencies():
  for name, pkg in _ALL_APP_PACKAGES.items():
    # print("name = " + name)
    if pkg[0] == "npm":
        npm_install(
            # Name this npm so that Bazel Label references look like @npm//package
            name = name,
            package_json = pkg[1],
            package_lock_json = pkg[2],
        )
    else:
        yarn_install(
            name = name,
            package_json = pkg[1],
            yarn_lock = pkg[2],
        )
