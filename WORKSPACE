# Bazel workspace created by @bazel/create 5.1.0

# Declares that this directory is the root of a Bazel workspace.
# See https://docs.bazel.build/versions/main/build-ref.html#workspace
workspace(
    # How this workspace would be referenced with absolute labels from another workspace
    name = "monorepo",
)

##############################################################
# Global config: versions
NODE_VERSION = "16.13.2"
YARN_VERSION = "1.22.17"

GOLANG_VERSION = "1.18.3"


##############################################################
# Bootstrapping rules.
# - Download rules
# - Fetch bootstrapping dependencies.

# By default, all external rules (like defined using "http_archive") are defined here.
load("//bazel:deps.bzl", "fetch_external_rules")

fetch_external_rules()

##############################################################
# Initialize external rules.

## NodeJS
load("@build_bazel_rules_nodejs//:repositories.bzl", "build_bazel_rules_nodejs_dependencies")

build_bazel_rules_nodejs_dependencies()

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories")

node_repositories(
    node_version=NODE_VERSION,
    yarn_version=YARN_VERSION,
)

load("//bazel/nodejs/tools:app_deps.bzl", "install_nodejs_app_dependencies")

install_nodejs_app_dependencies()

## Golang
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_register_toolchains(version=GOLANG_VERSION)

load("//bazel/golang:deps.bzl", "manual_go_deps")
load("//bazel/golang:generated_deps_from_go_mod.bzl", "fetch_go_mod_deps")

# gazelle:repository_macro bazel/golang/generated_deps_from_go_mod.bzl%fetch_go_mod_deps
fetch_go_mod_deps()

manual_go_deps()

go_rules_dependencies()

## Python
load("//bazel/python:deps.bzl", "install_all_pip_deps")

install_all_pip_deps()

###############################################################
# Bootstrap secondary rules.
# - Same as primary rules, load rules sources, fetch dependencies.
# The secondary rules depeneds on other rules.

## Protobuf
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

## Gazelle
# Dependencies:
# - Golang
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

# If you use WORKSPACE.bazel, use the following line instead of the bare gazelle_dependencies():
# gazelle_dependencies(go_repository_default_config = "@//:WORKSPACE.bazel")
gazelle_dependencies()
