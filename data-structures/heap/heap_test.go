package heap

import "testing"

func TestMinHeap(t *testing.T) {
	min := NewMinHeap()

	if !min.IsEmpty() {
		t.Error("Initial min heap should be empty")
	}

	min.Push(1, "First")
	min.Push(3, "Second")
	min.Push(20, "Third")
	min.Push(2, "Fourth")

	if min.Peek() != "First" {
		t.Errorf("Next value should be '%s' not '%s'", "First", min.Peek())
	}

	min.Pop()
	min.Pop()

	if min.Peek() != "Second" {
		t.Errorf("Next value should be '%s' not '%s'", "Second", min.Peek())
	}
}

func TestMaxHeap(t *testing.T) {
	max := NewMaxHeap()

	if !max.IsEmpty() {
		t.Error("Initial max heap should be empty")
	}

	max.Push(1, "First")
	max.Push(3, "Second")
	max.Push(20, "Third")
	max.Push(2, "Fourth")

	if max.Peek() != "Third" {
		t.Errorf("Next value should be '%s' not '%s'", "Third", max.Peek())
	}

	max.Pop()
	max.Pop()

	if max.Peek() != "Fourth" {
		t.Errorf("Next value should be '%s' not '%s'", "Second", max.Peek())
	}
}
