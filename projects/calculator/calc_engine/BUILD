load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "calc_engine",
    srcs = ["calc_engine.go"],
    importpath = "github.com/tedsilb/GoProjects/projects/calculator/calc_engine",
    visibility = ["//visibility:public"],
)

go_test(
    name = "calc_engine_test",
    srcs = ["calc_engine_test.go"],
    embed = [":calc_engine"],
)
