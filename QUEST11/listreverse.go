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

}*/

func ListReverse(l *List) {
	l2 := &List{}
	a := l.Head
	for a != nil {
		ListPushFront(l2, a.Data)
		a = a.Next
	}
	*l = *l2
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

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}*/
