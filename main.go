// bazel test //:greeter_test
// bazel build //:main
// bazel run main
// // package main

// // import "fmt"

// //create a channel in main function, pass any type of data into the channel,
// //create a go routine, pass this channel,
// //receive the data in same Go Routine, print the received data in go routine.

// // func main() {

// // 	ch := make(chan int)
// // 	go testRoutine(ch)
// // 	ch <- 10
// // 	// v := <-ch
// // 	// fmt.Println(v)

// // }

// // func testRoutine(ch chan int) {

// // 	fmt.Println(<-ch)

// // }

// package main

// import "fmt"

// // "time"

// // func display() {
// // 	for i := 0; i < 5; i++ {
// // 		fmt.Println("In display")
// // 	}
// // }

// type LinkedList struct {
// 	head *Node
// }
// type Node struct {
// 	data int
// 	next *Node
// }

// func (l *LinkedList) insert(node *Node) {
// 	if l.head == nil {
// 		l.head = node
// 		return
// 	}
// 	temp := l.head
// 	for l.head.next != nil {
// 		l.head = l.head.next
// 	}
// 	l.head.next = node
// 	l.head = temp

// 	return
// }

// func (l *LinkedList) Traverse() {

// 	for l.head != nil {
// 		fmt.Println(l.head.data)
// 		l.head = l.head.next
// 	}

// }

// func main() {

// 	l := &LinkedList{}
// 	node1 := &Node{data: 1}
// 	node2 := &Node{data: 2}
// 	node3 := &Node{data: 3}
// 	node4 := &Node{data: 4}
// 	l.insert(node1)
// 	l.insert(node2)
// 	l.insert(node3)
// 	l.insert(node4)
// 	l.Traverse()

// }

package main

import (
	"fmt"
	//"moduledep/greeter"
	net "net/http"

	"github.com/gin-gonic/gin"
)

type LinkedList struct {
	head *Node
}
type Node struct {
	data int
	next *Node
}

func main() {
	//fmt.Println(greeter.Greet())
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(net.StatusOK, "What a call bro !!!")
	})

	fmt.Println("Hello git")
	r.Run(":8080")
	fmt.Println("Hello git")
	// reverseLinkedList()
}

// git@github.com:alok-sonker/golangDS.git

func reverseLinkedList() {
	fmt.Println("Reverse a linked list")
	list := &LinkedList{}
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	node4 := &Node{data: 4}
	list.Insert(node1)
	list.Insert(node2)
	list.Insert(node3)
	list.Insert(node4)
	//fmt.Println(*&list.head.next.data)
	list.Traverse()
	//list.Reverse()
	recReverse(list.head)
	list.Traverse()
}

func (list *LinkedList) Insert(node *Node) {

	if list.head == nil {
		list.head = node
		return
	}
	temp := list.head
	for list.head.next != nil {
		list.head = list.head.next
	}
	list.head.next = node
	list.head = temp
}

func (l *LinkedList) setPointer() {

}

func (list LinkedList) Traverse() {

	for list.head != nil {
		fmt.Println(list.head.data)
		list.head = list.head.next
	}
}
func (list *LinkedList) Reverse() {
	temp := list.head
	//pre := &Node{}
	for list.head.next != nil {
		//fmt.Println(list.head.data)
		list.head = list.head.next
	}
	temp.next = nil
	list.head.next = temp

}
func recReverse(head *Node) {
	if head == nil {
		return
	}

	recReverse(head.next)
	fmt.Println(head.data)

}
