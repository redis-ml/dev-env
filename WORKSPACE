# Bazel workspace created by @bazel/create 5.1.0

# Declares that this directory is the root of a Bazel workspace.
# See https://docs.bazel.build/versions/main/build-ref.html#workspace
workspace(
    # How this workspace would be referenced with absolute labels from another workspace
    name = "monorepo",
)

##############################################################
# Load bootstrapping rules.

## NodeJS
load("//bazel/nodejs/tools:bazel_deps.bzl", "fetch_nodejs_dependencies")

fetch_nodejs_dependencies()

## Golang
load("//bazel/golang/internal:def.bzl", "bootstrap_golang_rules")

bootstrap_golang_rules()

## Python
load("//bazel/python:init.bzl", "bootstrap_python_rules")

bootstrap_python_rules()

##############################################################
# Load external rules.


##  NodeJS
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

