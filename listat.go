package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
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
func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; i < pos; i++ {
		if l != nil {
			l = l.Next
		} else {
			return nil
		}
	}
	return l
}

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}
