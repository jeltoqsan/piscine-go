package piscine

/*package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListSize(l *List) int {
	counter := 0
	node := l.Head
	for node != nil {
		counter++
		node = node.Next
	}
	return counter
}

/*func ListPushFront(l *List, data interface{}) {
	Node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = Node
		l.Tail = Node
		return
	}
	Node.Next = l.Head
	l.Head = Node
}

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}*/
