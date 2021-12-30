package tests

import (
	. "github.com/noodleslove/stack_n_queue/pkg/stack"
	"testing"
)

func TestStackConstructor(t *testing.T) {
	s := NewStack()
	if s == nil || !s.Empty() {
		t.Errorf("NewStack incorrect")
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack() // init Stack s
	for i := 0; i < 1000; i++ {
		s.Push(i) // push 0..999 to s
	}

	if s.Size() != 1000 { // test if stack size is correct
		t.Errorf("Push incorrect, got: %d, want: %d", s.Size(), 1000)
	}

	// test if stack elements are correct
	for i, x := s.Begin(), 999; i != s.End(); i, x = i.Next(), x-1 {
		if i.Elem != x {
			t.Errorf("Push incorrect, got: %d, want: %d", i.Elem, x)
		}
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack() // init Stack s
	for i := 0; i < 1000; i++ {
		s.Push(i) // push 0..999 to s
	}

	i := 999
	for !s.Empty() {
		val := s.Pop()
		if val != i {
			t.Errorf("Pop incorrect, got: %d, want: %d", val, i)
		}
		i--
	}
}
