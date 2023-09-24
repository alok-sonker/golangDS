// bazel test //:greeter_test
// bazel build //:main
// bazel run main

package main

func main() {
	list := InitList()
	list.Append(1, 0, 3, 5, 7)
	list.Traverse()
	list.Append(9, 10, 13, 15, 17)
	list.Traverse()
	list.Prepend(99, 100)
	list.Traverse()
}
