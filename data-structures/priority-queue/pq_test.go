package pq

import "testing"

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()

	if !pq.IsEmpty() {
		t.Error("Initial PQ should be empty")
	}

	pq.Enqueue(5, "First")
	pq.Enqueue(3, "Second")
	pq.Enqueue(10, "Third")

	if pq.Size() != 3 {
		t.Error("PQ should have 3 items in it")
	}

	if pq.Peek() != "Third" {
		t.Errorf("PQ's first item should be '%s' but found '%s'", "Third", pq.Peek())
	}

	first := pq.Dequeue()
	if first != "Third" {
		t.Errorf("PQ's first popped item should be '%s' but found '%s'", "Third", first)
	}

	second := pq.Dequeue()
	if second != "First" {
		t.Errorf("PQ's second popped item should be '%s' but found '%s'", "First", second)
	}

	third := pq.Dequeue()
	if third != "Second" {
		t.Errorf("PQ's third popped item should be '%s' but found '%s'", "Second", third)
	}

	if !pq.IsEmpty() {
		t.Error("PQ should be empty")
	}
}
