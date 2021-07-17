load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "calculator_lib",
    srcs = ["calculator.go"],
    importpath = "github.com/tedsilb/GoProjects/projects/calculator",
    visibility = ["//visibility:private"],
    deps = [
        "//projects/calculator/calc_engine",
    ],
)

go_binary(
    name = "calculator",
    embed = [":calculator_lib"],
    visibility = ["//visibility:public"],
)
