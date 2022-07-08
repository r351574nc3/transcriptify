load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

gazelle(
    name = "gazelle",
    prefix = "github.com/r351574nc3/transcriptify",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

genrule(
    name = "app",
    srcs = [
        "//cmd:cli",
    ],
    outs = [
        "transcriptify",
    ],
    cmd = "cp ./bazel-out/*/bin/cmd/*/cli \"$@\"",
    executable = 1,
    local = 1,
)

