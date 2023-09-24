// bazel test //:greeter_test
// bazel build //:main
// bazel run main

package main

import "fmt"

func main() {
	list := InitList()
	list.Append(1, 0, 3, 5, 7)
	list.Traverse()
	list.Append(9, 10, 13, 15, 17)
	fmt.Println("Testing mergeconflict")
	list.Traverse()
	list.Prepend(99, 100)
	list.Traverse()
}
