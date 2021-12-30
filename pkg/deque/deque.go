package deque

import (
	. "github.com/noodleslove/stack_n_queue/pkg/node"
)

type Deque struct {
	head *Node
	tail *Node
	size uint32
}

func NewDeque() *Deque {
	p := Deque{head: nil, tail: nil, size: 0}
	return &p
}

func NewDequeFromList(elems []interface{}) *Deque {
	d := NewDeque()
	for _, elem := range elems {
		d.PushBack(elem)
	}
	return d
}

func (d *Deque) Empty() bool {
	return d.size == 0
}

/**
 * Precondition: size() > 0
 * Postcondition: The return value is the front item of the deque (but this
 *      item is not removed from the deque).
 *
 */
func (d *Deque) Front() interface{} {
	if d.Empty() {
		panic("Can't pop empty deque")
	}

	return At(d.head, 0)
}

/**
 * Precondition: size() > 0
 * Postcondition: The return value is the tail item of the deque (but this
 *      item is not removed from the deque).
 *
 */
func (d *Deque) Back() interface{} {
	if d.Empty() {
		panic("Can't pop empty deque")
	}

	return At(d.head, d.size-1)
}

func (d *Deque) Begin() *Node {
	if d.Empty() {
		return nil
	}

	return d.head
}

func (d *Deque) End() *Node {
	return nil
}

/**
 * Postcondition: The return value is the number of elements in the deque.
 *
 */
func (d *Deque) Size() uint32 {
	return d.size
}

/**
 * Postcondition: Insert the given item at the beginning of the deque
 *
 */
func (d *Deque) PushFront(elem interface{}) {
	// update deque size
	d.size++

	// init deque if it is empty
	if d.head == nil {
		d.tail = InsertHead(&d.head, elem)
		return
	}

	// not empty, insert item to head and update head
	InsertHead(&d.head, elem)
}

func (d *Deque) PushBack(elem interface{}) {
	// update deque size
	d.size++
	// init deque if it is empty
	if d.tail == nil {
		d.tail = InsertHead(&d.head, elem)
		return
	}

	// not empty, insert item to tail and update tail
	d.tail = InsertAfter(&d.head, d.tail, elem)
}

/**
 * Precondition: size() > 0
 * Postcondition: The top item of the deque has been removed. The return value
 *      is the _item value of the removed node.
 *
 */
func (d *Deque) PopFront() interface{} {
	if d.Empty() {
		panic("Can't pop empty deque")
	}

	// update size
	d.size--
	// delete and update head node
	// return _item of it
	return DeleteNode(&d.head, d.head)
}

/**
 * Precondition: size() > 0
 * Postcondition: The last item of the deque has been removed. The return value
 *      is the _item value of the removed node.
 *
 */
func (d *Deque) PopBack() interface{} {
	if d.Empty() {
		panic("Can't pop empty deque")
	}

	// update size
	d.size--
	// delete and update tail node
	// return _item of it
	delete_this := d.tail
	d.tail = PreviousNode(d.tail)
	return DeleteNode(&d.head, delete_this)
}
