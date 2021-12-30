package stack

import (
	. "github.com/noodleslove/stack_n_queue/pkg/node"
)

type Stack struct {
	top  *Node
	size uint32
}

func NewStack() *Stack {
	p := Stack{top: nil, size: 0}
	return &p
}

/**
 * Postcondition: The return value is true if stack object has no element,
 *      else return false.
 *
 */
func (s *Stack) Empty() bool {
	return s.size == 0
}

/**
 * Postcondition: The return value is the number of elements in the stack.
 *
 */
func (s *Stack) Size() uint32 {
	return s.size
}

/**
 * Precondition: size() > 0
 * Postcondition: The return value is the front item of the stack (but this
 *      item is not removed from the queue).
 *
 */
func (s *Stack) Top() interface{} {
	if s.Empty() {
		panic("Can't Top empty stack")
	}

	return s.top.Elem
}

/**
 * Postcondition: The return value is an iterator to the top of the stack.
 *
 */
func (s *Stack) Begin() *Node {
	if s.Empty() {
		return nil
	}

	return s.top
}

/**
 * Postcondition: The return value is an iterator to nullptr.
 *
 */
func (s *Stack) End() *Node {
	return nil
}

/**
 * Postcondition: Print all the elements in the queue to standard output.
 *
 */
func (s *Stack) Display() {
	PrintList(s.top)
}

/**
 * Postcondition: A new element has been inserted at the top of the stack.
 *
 * @param item The item to be inserted.
 */
func (s *Stack) Push(elem interface{}) {
	s.size++
	InsertHead(&s.top, elem)
}

/**
 * Precondition: size() > 0
 * Postcondition: The top item of the stack has been removed. The return value
 *      is the _item value of the removed node.
 *
 */
func (s *Stack) Pop() interface{} {
	if s.Empty() {
		panic("Can't Pop empty stack")
	}

	s.size--
	return DeleteNode(&s.top, s.top)
}
