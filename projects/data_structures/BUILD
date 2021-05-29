load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "data_structures_lib",
    srcs = ["data_structures.go"],
    importpath = "github.com/tedsilb/GoProjects/projects/data_structures",
    visibility = ["//visibility:private"],
    deps = [
        "//projects/data_structures/doubly_linked_list",
        "//projects/data_structures/linked_list",
    ],
)

go_binary(
    name = "data_structures",
    embed = [":data_structures_lib"],
    visibility = ["//visibility:public"],
)
