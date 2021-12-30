package queue

import (
	. "fake.com/tutorial/node"
)

type Queue struct {
	head *Node
	tail *Node
	size uint32
}

func NewQueue() *Queue {
	p := Queue{head: nil, tail: nil, size: 0}
	return &p
}

/**
 * Postcondition: The return value is true is queue object is empty, else return
 *       false.
 *
 */
func (q *Queue) Empty() bool {
	return q.size == 0
}

/**
 * Postcondition: The return value is the number of elements in the queue.
 *
 */
func (q *Queue) Size() uint32 {
	return q.size
}

/**
 * Precondition: size() > 0
 * Postcondition: The return value is the front item of the queue (but this
 *      item is not removed from the queue).
 *
 */
func (q *Queue) Front() interface{} {
	if q.Empty() {
		panic("Can't front empty queue")
	}

	return At(q.head, 0)
}

/**
 * Precondition: size() > 0
 * Postcondition: The return value is the tail item of the queue (but this
 *      item is not removed from the queue).
 *
 */
func (q *Queue) Back() interface{} {
	if q.Empty() {
		panic("Can't back empty queue")
	}

	return At(q.head, q.size-1)
}

/**
 * Postcondition: The return value is an iterator to the head of the queue.
 *
 */
func (q *Queue) Begin() *Node {
	if q.Empty() {
		return nil
	}

	return q.head
}

/**
 * Postcondition: The return value is an iterator to nullptr.
 *
 */
func (q *Queue) End() *Node {
	return nil
}

/**
 * Postcondition: Print all the elements in the queue to standard output.
 *
 */
func (q *Queue) Display() {
	PrintList(q.head)
}

/**
 * Postcondition: A new element has been inserted at the tail of the
 *      queue.
 *
 * @param item The item to be inserted.
 */
func (q *Queue) Push(elem interface{}) {
	// update queue size
	q.size++
	// init queue when it is empty
	if q.tail == nil {
		q.tail = InsertHead(&q.head, elem)
		return
	}

	// insert item to tail and update _tail
	q.tail = InsertAfter(&q.head, q.tail, elem)
}

/**
 * Precondition: size() > 0
 * Postcondition: The top item of the queue has been removed. The return value
 *      is the _item value of the removed node.
 *
 */
func (q *Queue) Pop() interface{} {
	if q.Empty() {
		panic("Can't pop empty queue")
	}
	// update size
	q.size--
	// delete and update head node
	// return _item of it
	return DeleteNode(&q.head, q.head)
}
