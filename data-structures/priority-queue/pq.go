package pq

import heap "../heap"

// PriorityQueue uses a Max Heap to handle all operations.
type PriorityQueue struct {
	items *heap.MaxHeap
}

// NewPriorityQueue creates a new instance of a priority queue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{items: heap.NewMaxHeap()}
}

// Enqueue pushes a new value onto the back of the priority queue.
func (q *PriorityQueue) Enqueue(pr int, val interface{}) {
	q.items.Push(pr, val)
}

// Peek returns the next value without popping it off.
func (q *PriorityQueue) Peek() interface{} {
	return q.items.Peek()
}

// Dequeue gets the largest value from the heap and return it.
func (q *PriorityQueue) Dequeue() interface{} {
	return q.items.Pop()
}

// IsEmpty checks if there are any items currently in the priority queue.
func (q *PriorityQueue) IsEmpty() bool {
	return q.items.IsEmpty()
}

// Size checks the length of the priority queue.
func (q *PriorityQueue) Size() int {
	return q.items.Cardinality()
}
