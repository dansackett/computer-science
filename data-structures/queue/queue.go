package queue

import list "../linked-list"

// Queue holds the items.
type Queue struct {
	Items *list.List
}

// NewQueue returns an empty queue.
func NewQueue() *Queue {
	return &Queue{Items: list.New()}
}

// Enqueue pushes a new value onto the back of the queue.
func (q *Queue) Enqueue(val interface{}) {
	q.Items.Prepend(val)
}

// Dequeue pops the last value of the queue off removing it from the
// underlying list and returning the value.
func (q *Queue) Dequeue() interface{} {
	return q.Items.Pop().Value
}

// IsEmpty checks if there are any items currently in the queue.
func (q *Queue) IsEmpty() bool {
	if q.Size() > 0 {
		return false
	}
	return true
}

// Size checks the length of the queue
func (q *Queue) Size() int {
	return q.Items.Size()
}
