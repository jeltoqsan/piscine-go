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

func ListLast(l *List) interface{} {
	node := l.Head
	if l.Head == nil {
		return nil
	} else {
		for l.Head != nil {
			node = l.Head
			l.Head = l.Head.Next
		}
		return node.Data
	}
}

/*func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		return
	}
	l.Tail = &NodeL{Data: data}
	n := l.Head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &NodeL{Data: data}

}
func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}*/
