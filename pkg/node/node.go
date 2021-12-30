package node

import "fmt"

// ----------------------------------------
// `Node` implementation
// ----------------------------------------

type Node struct {
	Elem interface{}
	next *Node
	prev *Node
}

func NewNode(val interface{}, next *Node, prev *Node) *Node {
	p := Node{Elem: val, next: next, prev: prev}
	return &p
}

func (n *Node) Next() *Node {
	if n != nil {
		return n.next
	}

	return nil
}

func (n *Node) Prev() *Node {
	if n != nil {
		return n.prev
	}

	return nil
}

func (n *Node) Display() {
	if n != nil {
		fmt.Printf("<-[%+v]->", n.Elem)
	}
}

// ----------------------------------------
// `List` implementation
// ----------------------------------------

type List struct {
	head *Node
	size uint32
}

func NewList() *List {
	p := List{head: nil, size: 0}
	return &p
}

func (l *List) Push(elem interface{}) {
	newHead := &Node{elem, l.head, nil}

	if l.head != nil {
		l.head.prev = newHead
	}

	l.head = newHead
	l.size += 1
}

func (l *List) Pop() *interface{} {
	if l.size > 0 {
		ret := l.head.Elem
		newHead := l.head.next

		l.head.next = nil
		l.head = newHead
		l.head.prev = nil
		l.size -= 1
		return &ret
	}

	return nil
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		list.Display()
		list = list.next
	}
	fmt.Println("|||")
}
