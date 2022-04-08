# Bazel workspace created by @bazel/create 5.1.0

# Declares that this directory is the root of a Bazel workspace.
# See https://docs.bazel.build/versions/main/build-ref.html#workspace
workspace(
    # How this workspace would be referenced with absolute labels from another workspace
    name = "monorepo",
)

##############################################################
# Bootstrapping rules.
# - Download rules
# - Fetch bootstrapping dependencies.

# By default, all external rules (like defined using "http_archive") are defined here.
load("//:bazel/deps.bzl", "fetch_external_rules")

fetch_external_rules()

##############################################################
# Initialize external rules.

## NodeJS
load("@build_bazel_rules_nodejs//:repositories.bzl", "build_bazel_rules_nodejs_dependencies")

build_bazel_rules_nodejs_dependencies()

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories")

node_repositories(
    node_version = "16.13.2",
    yarn_version = "1.22.17",
)

load("//bazel/nodejs/tools:app_deps.bzl", "install_nodejs_app_dependencies")

install_nodejs_app_dependencies()

## Golang
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.17.6")

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
