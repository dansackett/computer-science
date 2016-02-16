package stack

import list "../linked-list"

// Stack holds the items.
type Stack struct {
	Items *list.List
}

// NewStack returns an empty stack.
func NewStack() *Stack {
	return &Stack{Items: list.New()}
}

// Push places a new value on the end of the stack
func (q *Stack) Push(val interface{}) {
	q.Items.Append(val)
}

// Pop removes the value at the end of the stack
func (q *Stack) Pop() interface{} {
	return q.Items.Pop().Value
}

// IsEmpty checks if there are any items currently in the stack.
func (q *Stack) IsEmpty() bool {
	if q.Size() > 0 {
		return false
	}
	return true
}

// Size checks the length of the stack
func (q *Stack) Size() int {
	return q.Items.Size()
}
