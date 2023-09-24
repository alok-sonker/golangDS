package main

import (
	"fmt"
)

type listOperations interface {
	Append(...int32)
	Prepend(...int32)
	Traverse()
	Delete(int32)
	Update(int32)
	Search(int32)
	Sort()
	Reverse()
}

func InitList() *LinkedList {
	return &LinkedList{}
}
func createNode(num int32) *Node {
	return &Node{data: num}
}

func (list *LinkedList) isEmptykHead() bool {
	return list.head == nil
}

func (list *LinkedList) Prepend(node ...int32) {
	if len(node) == 0 {
		return
	}
	tempList := InitList()
	tempNode := &Node{}
	for _, n := range node {
		newlyCreatedNode := createNode(n)
		if tempList.head == nil {
			tempList.head = newlyCreatedNode
			tempNode = tempList.head
			continue
		}
		tempList.head.next = newlyCreatedNode
		tempList.head = tempList.head.next
	}
	tempList.head.next = list.head
	list.head = tempNode
}

func (list *LinkedList) Append(node ...int32) {
	if len(node) == 0 {
		return
	}
	tempList := InitList()
	tempNode := &Node{}
	for _, n := range node {
		newlyCreatedNode := createNode(n)
		if tempList.head == nil {
			tempList.head = newlyCreatedNode
			tempNode = tempList.head
			continue
		}
		tempList.head.next = newlyCreatedNode
		tempList.head = tempList.head.next
	}
	tempList.head = tempNode
	if list.isEmptykHead() {
		list.head = tempList.head
		return
	}
	mainListHead := list.head
	for list.head.next != nil {
		list.head = list.head.next
	}
	list.head.next = tempList.head
	list.head = mainListHead
}

func (list LinkedList) Traverse() {
	for list.head != nil {
		fmt.Printf("%d ", list.head.data)
		list.head = list.head.next
	}
	fmt.Println()
}
func (list *LinkedList) Reverse() {
	temp := list.head
	for list.head.next != nil {
		list.head = list.head.next
	}
	temp.next = nil
	list.head.next = temp

}
func recursiveReverse(head *Node) {
	if head == nil {
		return
	}
	recursiveReverse(head.next)
	fmt.Println(head.data)

}
