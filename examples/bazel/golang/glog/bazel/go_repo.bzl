load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_deps():
    go_repository(
        name = "com_github_spyzhov_ajson",
        importpath = "github.com/spyzhov/ajson",
        sum = "h1:sFXyMbi4Y/BKjrsfkUZHSjA2JM1184enheSjjoT/zCc=",
        version = "v0.8.0",
    )
