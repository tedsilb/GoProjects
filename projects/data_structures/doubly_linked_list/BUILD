load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "doubly_linked_list",
    srcs = ["doubly_linked_list.go"],
    importpath = "github.com/tedsilb/GoProjects/projects/data_structures/doubly_linked_list",
    visibility = ["//visibility:public"],
)

go_test(
    name = "doubly_linked_list_test",
    srcs = ["doubly_linked_list_test.go"],
    embed = [":doubly_linked_list"],
)
