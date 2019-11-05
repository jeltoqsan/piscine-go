package piscine

func ListPushFront(l *List, data interface{}) {
	Node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = Node
		l.Tail = Node
		return
	}
	Node.Next = l.Head
	l.Head = Node
}
