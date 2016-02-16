package deque

import list "../linked-list"

// Deque holds the items.
type Deque struct {
	Items *list.List
}

// NewDeque returns an empty deque.
func NewDeque() *Deque {
	return &Deque{Items: list.New()}
}

// Append pushes an item to the back of the deque.
func (q *Deque) Append(val interface{}) {
	q.Items.Append(val)
}

// Prepend pushes an item to the front of the deque.
func (q *Deque) Prepend(val interface{}) {
	q.Items.Prepend(val)
}

// Pop pulls an item off the back of the deque.
func (q *Deque) Pop() interface{} {
	return q.Items.Pop().Value
}

// PopLeft pulls and item off the front of the deque.
func (q *Deque) PopLeft() interface{} {
	return q.Items.PopLeft().Value
}

// IsEmpty checks if there are any items currently in the deque.
func (q *Deque) IsEmpty() bool {
	if q.Size() > 0 {
		return false
	}
	return true
}

// Size checks the length of the deque.
func (q *Deque) Size() int {
	return q.Items.Size()
}
