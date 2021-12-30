package tests

import (
	. "fake.com/tutorial/queue"
	"testing"
)

func TestQueueConstructor(t *testing.T) {
	q := NewQueue()
	if q == nil || !q.Empty() {
		t.Errorf("NewQueue incorrect")
	}
}

func TestQueuePush(t *testing.T) {
	q := NewQueue() // init Queue q
	for i := 0; i < 1000; i++ {
		q.Push(i) // push 0..999 to q
	}

	if q.Size() != 1000 { // test if queue size is correct
		t.Errorf("Push incorrect, got: %d, want: %d", q.Size(), 1000)
	}

	// test if queue elements are correct
	for i, x := q.Begin(), 0; i != q.End(); i, x = i.Next(), x+1 {
		if i.Elem != x {
			t.Errorf("Push incorrect, got: %d, want: %d", i.Elem, x)
		}
	}
}

func TestQueuePop(t *testing.T) {
	q := NewQueue() // init Queue q
	for i := 0; i < 1000; i++ {
		q.Push(i) // push 0..999 to q
	}

	i := 0
	for !q.Empty() {
		val := q.Pop()
		if val != i {
			t.Errorf("Pop incorrect, got: %d, want: %d", val, i)
		}
		i++
	}
}
