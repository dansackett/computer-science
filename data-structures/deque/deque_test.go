package deque

import "testing"

func TestInitialDequeIsEmpty(t *testing.T) {
	deque := NewDeque()

	if !deque.IsEmpty() {
		t.Errorf("Deque is supposed to be empty")
	}
}

func TestDequeWorks(t *testing.T) {
	deque := NewDeque()
	// This should give us [8, 5, 1]
	deque.Append(5)
	deque.Prepend(8)
	deque.Append(1)

	if deque.Size() != 3 {
		t.Errorf("Deque should have 3 items in it, it has %d", deque.Size())
	}

	first := deque.Pop()

	if deque.Size() != 2 {
		t.Errorf("Deque should have 2 items in it, it has %d", deque.Size())
	}

	if first != 1 {
		t.Errorf("First value should be 1, it returned %d", first)
	}

	second := deque.PopLeft()
	third := deque.Pop()

	if !deque.IsEmpty() {
		t.Errorf("Deque is supposed to be empty")
	}

	if second != 8 {
		t.Errorf("Second value should be 8, it returned %d", second)
	}

	if third != 5 {
		t.Errorf("Third value should be 5, it returned %d", third)
	}
}
