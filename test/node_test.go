package tests

import (
	. "fake.com/tutorial/node"
	"testing"
)

func TestNode(t *testing.T) {
	n := NewNode(1, nil, nil)
	if n == nil || n.Elem != 1 {
		t.Errorf("Node initialize incorrect, got: %+v, want: %+v", n, n)
	}
}
