http_archive(
    name = "io_bazel_rules_go",
    sha256 = "b7a62250a3a73277ade0ce306d22f122365b513f5402222403e507f2f997d421",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.16.3/rules_go-0.16.3.tar.gz",
)

git_repository(
    name = "bazel_gazelle",
    commit = "af1d52d3fca8ab66fa49e0d11dd70d03a5a1d770",
    remote = "https://github.com/yext/bazel-gazelle",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "org_golang_x_net",
    commit = "ab5485076ff3407ad2d02db054635913f017b0ed",
    importpath = "golang.org/x/net",
    build_file_proto_mode = "disable",
    build_file_name = "BUILD.bazel",
)
