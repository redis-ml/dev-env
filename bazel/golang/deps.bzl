load("@bazel_gazelle//:deps.bzl", "go_repository")

def manual_go_deps():
    # go-containerregistry is a dependency that typically comes from
    # @io_bazel_rules_docker//repositories:deps.bzl, but we manually override it
    # here to get a more recent version that has it's own dependencies
    # simplified.
    #
    # Commit: 628a2ff5f006eca399a316a66cc714106fcb3943
    # Date: 2021-07-08 17:54:56 +0000 UTC
    # URL: https://github.com/google/go-containerregistry/commit/628a2ff5f006eca399a316a66cc714106fcb3943
    #
    # Document released images in release notes (#1069)
    # Size: 3991527 (4.0 MB)
    go_repository(
        name = "com_github_google_go_containerregistry",
        importpath = "github.com/google/go-containerregistry",
        commit = "628a2ff5f006eca399a316a66cc714106fcb3943",
        build_directives = [
            # the k8schain package is not used.  Gazelle and go modules are
            # confused by it:
            # https://github.com/vdemeester/k8s-pkg-credentialprovider#k8siokubernetespkgcredentialprovider-temporary-fork
            "gazelle:exclude pkg/authn/k8schain/**/*",
        ],
    )
