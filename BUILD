load("@io_bazel_rules_go//go:def.bzl","go_binary","go_library","go_test")

go_binary(
	name="main",
	srcs=["main.go"],
	deps=[":greeter"],
)
go_library(
	name="greeter",
	importpath="dep/greeter",
	srcs=["greeter/greeter.go"],

	
)
go_test(
	name="greeter_test",
	srcs=["greeter/greeter_test.go"],
	embed=[":greeter"],
)