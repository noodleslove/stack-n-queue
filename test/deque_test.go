package tests

import (
	. "fake.com/tutorial/deque"
	"testing"
)

func TestDequeConstructor(t *testing.T) {
	ans := make([]interface{}, 1000)
	for i := range ans {
		ans[i] = i + 1
	}

	d := NewDequeFromList(ans)
	itr := d.Begin()
	for _, elem := range ans {
		if elem != itr.Elem {
			t.Errorf("NewDequeFromList incorrect")
		}
		itr = itr.Next()
	}
}

func TestDequePush(t *testing.T) {
	d := NewDeque()

	for i := 1; i < 10; i++ {
		d.PushFront(i * 10)

		if d.Front() != i*10 {
			t.Errorf("PushFront incorrect, got: %d, want: %d", d.Front(), i*10)
		}
	}

	for i := 0; i > -10; i-- {
		d.PushBack(i * 10)

		if d.Back() != i*10 {
			t.Errorf("PushFront incorrect, got: %d, want: %d", d.Back(), i*10)
		}
	}
}

func TestDequePop(t *testing.T) {
	d := NewDeque()

	for i := 1; i <= 10; i++ {
		d.PushBack(i * 10)
	}

	for i := 1; i <= 5; i++ {
		val := d.PopFront()

		if val != i*10 {
			t.Errorf("PopFront incorrect, got: %d, want: %d", val, i*10)
		}
	}

	for i := 10; i >= 6; i-- {
		val := d.PopBack()

		if val != i*10 {
			t.Errorf("PopBack incorrect, got: %d, want: %d", val, i*10)
		}
	}
}

func TestDequeEmpty(t *testing.T) {
	d := NewDeque()

	if !d.Empty() {
		t.Errorf("Empty incorrect, got: %t, want:%t", d.Empty(), true)
	}

	d.PushBack(0)
	if d.Empty() {
		t.Errorf("Empty incorrect, got: %t, want: %t", d.Empty(), false)
	}

	d.PopFront()
	if !d.Empty() {
		t.Errorf("Empty incorrect, got: %t, want:%t", d.Empty(), true)
	}
}

func TestDequeIterator(t *testing.T) {
	ans := make([]interface{}, 1000)
	for i := range ans {
		ans[i] = i + 1
	}

	d := NewDequeFromList(ans)
	for i, x := d.Begin(), 0; i != nil; i, x = i.Next(), x+1 {
		if i.Elem != x+1 {
			t.Errorf("DequeIterator incorrect")
		}
	}
}
