package main

type ds interface {
	array()
	slice()
	stack()
	queue()
	hashmap()
	linkedlist()
	heap()
	tree()
	graph()
}

func (ds DS) array() {

}
func (ds DS) stack() {

}
func (ds DS) queue() {

}
func (ds DS) slice() {

}
func (ds DS) hashmap() {

}
func (ds DS) linkedlist() {
	list := InitList()
	list.Append(1, 0, 3, 5, 7)
	list.Traverse()
	list.Append(9, 10, 13, 15, 17)
	list.Traverse()
	list.Prepend(99, 100)
	list.Traverse()
}
func (ds DS) heap() {

}
func (ds DS) tree() {

}
func (ds DS) graph() {

}
