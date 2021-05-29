load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "linked_list",
    srcs = ["linked_list.go"],
    importpath = "github.com/tedsilb/GoProjects/projects/data_structures/linked_list",
    visibility = ["//visibility:public"],
)

go_test(
    name = "linked_list_test",
    srcs = ["linked_list_test.go"],
    embed = [":linked_list"],
)
