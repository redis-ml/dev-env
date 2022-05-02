
PATH := $(bazel run @go_sdk//:bin/go -- env GOROOT)/bin:$(PATH)

yarn:
	bazel run @yarn//:bin/yarn --

go:
	bazel run @go_sdk//:bin/go --
