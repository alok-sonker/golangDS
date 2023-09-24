package main

// LinkedList
type LinkedList struct {
	len  int32
	head *Node
}
type Node struct {
	data int32
	next *Node
}
