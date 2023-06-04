# Golang example


## Usage

1. To run the main function: `bazel run //glog`

2. To update dependencies:

  1. Add new dependnecies using Golang: `go get ....`, or `bazel run @go_sdk//:bin/go -- get ...`.

  1. Update BUILD files using Gazelle: `bazel run //:gazel-update-repos`

      1. the dependencies will be written to `bazel/go_deps.bzl`, which is invoked in WORKSPACE.

