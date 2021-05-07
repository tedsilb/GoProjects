load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "descriptive_stats",
    embed = [":descriptive_stats_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "descriptive_stats_lib",
    srcs = ["descriptive_stats.go"],
    importpath = "github.com/tedsilb/GoProjects/projects/descriptive_stats",
    visibility = ["//visibility:private"],
)
