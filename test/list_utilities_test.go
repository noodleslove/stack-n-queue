package tests

import (
	. "fake.com/tutorial/node"
	"testing"
)

func TestInsertHead(t *testing.T) {
	n := NewNode(1, nil, nil)
	n2 := InsertHead(&n, 2)
	if n2 != n {
		t.Errorf("InsertHead incorrect: got: %+v, %+v", n, n2)
	}
}

func TestInsertBefore(t *testing.T) {
	n := NewNode(1, nil, nil)
	InsertHead(&n, 2)
	n1 := InsertHead(&n, 3)
	InsertHead(&n, 4)
	InsertBefore(&n, n1, 0)
	ans := []int{4, 0, 3, 2, 1}
	for itr, idx := n, 0; itr != nil; itr, idx = itr.Next(), idx+1 {
		if itr.Elem != ans[idx] {
			t.Errorf("InsertBefore incorrect")
		}
	}

	var n2 *Node = nil
	InsertBefore(&n2, n, "a")
	InsertBefore(&n2, n2, "b")
	ans2 := []string{"b", "a"}
	for itr, idx := n2, 0; itr != nil; itr, idx = itr.Next(), idx+1 {
		if itr.Elem != ans2[idx] {
			t.Errorf("InsertBefore incorrect")
		}
	}
}

func TestDeleteNode(t *testing.T) {
	n := NewNode(1, nil, nil)
	n1 := InsertHead(&n, 2)
	InsertHead(&n, 3)
	InsertHead(&n, 4)

	x := DeleteNode(&n, n)
	if x != 4 {
		t.Errorf("DeleteNode incorrect, got: %d, want: %d", x, 4)
	}

	x = DeleteNode(&n, n1)
	if x != 2 {
		t.Errorf("DeleteNode incorrect, got: %d, want: %d", x, 2)
	}
}

func TestInsertAfter(t *testing.T) {
	n := NewNode(1, nil, nil)
	InsertHead(&n, 2)
	n1 := InsertHead(&n, 3)
	InsertHead(&n, 4)

	InsertAfter(&n, n1, 0)
	ans := []int{4, 3, 0, 2, 1}
	for itr, idx := n, 0; itr != nil; itr, idx = itr.Next(), idx+1 {
		if itr.Elem != ans[idx] {
			t.Errorf("InsertBefore incorrect")
		}
	}

	var n2 *Node = nil
	InsertAfter(&n2, nil, "a")
	InsertAfter(&n2, n2, "b")
	ans2 := []string{"a", "b"}
	for itr, idx := n2, 0; itr != nil; itr, idx = itr.Next(), idx+1 {
		if itr.Elem != ans2[idx] {
			t.Errorf("InsertBefore incorrect")
		}
	}
}
