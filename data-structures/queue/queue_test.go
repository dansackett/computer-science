package queue

import "testing"

func TestInitialQueueIsEmpty(t *testing.T) {
	queue := NewQueue()

	if !queue.IsEmpty() {
		t.Errorf("Queue is supposed to be empty")
	}
}

func TestQueueWorks(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(5)
	queue.Enqueue(8)
	queue.Enqueue(1)

	if queue.Size() != 3 {
		t.Errorf("Queue should have 3 items in it, it has %d", queue.Size())
	}

	first := queue.Dequeue()

	if queue.Size() != 2 {
		t.Errorf("Queue should have 2 items in it, it has %d", queue.Size())
	}

	if first != 5 {
		t.Errorf("First value should be 5, it returned %d", first)
	}

	second := queue.Dequeue()
	third := queue.Dequeue()

	if !queue.IsEmpty() {
		t.Errorf("Queue is supposed to be empty")
	}

	if second != 8 {
		t.Errorf("Second value should be 8, it returned %d", second)
	}

	if third != 1 {
		t.Errorf("Third value should be 1, it returned %d", third)
	}
}
