package node

import "fmt"

/**
 * Postcondition: Print the list to standard output.
 *
 * @param head node ptr
 */
func PrintList(head *Node) {
	// base case
	if head == nil {
		fmt.Println("|||")
		return
	}
	if head.prev == nil {
		fmt.Printf("Head->") // print "Head->" before head
	}

	head.Display()
	fmt.Printf(" ")
	PrintList(head.next)
}

/**
 * Postcondition: The return value is the node with match key. (if there is no
 *      match node, return value is nullptr)
 *
 * @param head node ptr
 * @param key to be searched
 * @return ptr to key or nullptr
 */
func SearchList(head *Node, key interface{}) *Node {
	// empty list return nullptr
	if head == nil {
		return nil
	}

	// search process
	curr := head
	for curr != nil {
		if curr.Elem == key {
			return curr // found key, return match node
		}
		curr = curr.next
	}

	// cannot find key, return nullptr
	return nil
}

/**
 * Postcondition: Insert a new node at the beginning of the list. The return
 *      value is the new head node of the list.
 *
 * @param head node ptr
 * @param value to be inserted
 * @return ptr to new head node
 */
func InsertHead(head **Node, insert_this interface{}) *Node {
	// init new head
	newHead := &Node{Elem: insert_this, next: *head, prev: nil}

	// check if head is `nil`, if not modify Prev reference
	if *head != nil {
		(*head).prev = newHead
	}
	// change head to new head
	*head = newHead
	// return new head
	return newHead
}

/**
 * Postcondition: Insert a new node after after_this node with insert_this
 *      value. (if after_this is nullptr, insert at the beginning) The return
 *      value is the new head of the list.
 *
 * @param head node ptr
 * @param after_this insert position node
 * @param insert_this value to be inserted
 * @return node<T>* ptr to inserted node
 */
func InsertAfter(
	head **Node,
	after_this *Node,
	insert_this interface{},
) *Node {
	// check if after_this is `nil`, if so insert at head
	if head == nil {
		return InsertHead(head, insert_this)
	}

	// search for after_this node, if find then insert
	curr := *head
	for curr != nil {
		if curr == after_this {
			if curr.next != nil {
				curr.next.prev = &Node{
					Elem: insert_this,
					next: curr.next,
					prev: curr,
				}
				curr.next = curr.next.prev
			} else {
				curr.next = &Node{Elem: insert_this, next: nil, prev: curr}
			}

			return curr.next
		}

		curr = curr.next
	}

	// cannot find after_this node, insert at head
	return InsertHead(head, insert_this)
}

/**
 * Postcondition: Insert a new node after before_this node with insert_this
 *      value. (if before_this is nullptr, insert at the beginning) The return
 *      value is the newly node inserted node.
 *
 * @param head node ptr
 * @param before_this insert position ptr
 * @param insert_this value to be inserted
 * @return node<T>* ptr to inserted node
 */
func InsertBefore(
	head **Node,
	before_this *Node,
	insert_this interface{},
) *Node {
	// check if before_this is nullptr, if so insert at head
	if before_this == nil {
		return InsertHead(head, insert_this)
	}

	// search for before_this node, if find then insert
	curr := *head
	for curr != nil {
		if curr.next == before_this {
			curr.next.prev = &Node{
				Elem: insert_this,
				next: curr.next,
				prev: curr,
			}
			curr.next = curr.next.prev
			return curr.next
		}

		curr = curr.next
	}

	// cannto find before_this, insert at head
	return InsertHead(head, insert_this)
}

func PreviousNode(prev_to_this *Node) *Node {
	// return _prev node of current node
	return prev_to_this.prev
}

func DeleteNode(
	head **Node,
	delete_this *Node,
) interface{} {
	// init ptrs
	var ret interface{}
	curr := *head

	// search for delete_this node, then delete
	for curr != nil {
		if curr == delete_this {
			ret = curr.Elem
			if curr.prev != nil && curr.next != nil {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
			} else if curr.prev != nil {
				curr.prev.next = nil
			} else if curr.next != nil {
				curr.next.prev = nil
				*head = curr.next
			} else {
				*head = nil
			}
			break
		}
		curr = curr.next
	}

	return ret
}

func At(head *Node, pos uint32) interface{} {
	curr := head

	// searching pos
	for i := 0; i < int(pos); i++ {
		curr = curr.next
	}
	return curr.Elem
}
