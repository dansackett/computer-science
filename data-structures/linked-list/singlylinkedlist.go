package linkedlist

import ()

// SingleNode is a node in a SinglyLinkedList
type SingleNode struct {
	// Value stored for the node
	Value interface{}
	// Next item is a SingleNode which follows this one
	Next *SingleNode
}

// Returns an initialized SingleNode
func NewSingleNode(v interface{}) *SingleNode {
	return &SingleNode{Value: v}
}

// Returns the item that follows a node
func (n *SingleNode) Next() *SingleNode {
	if n.Next {
		return n.Next
	}
	return nil
}

// SingleList is an instance of a SinglyLinkedList
type SingleList struct {
	// Root node or the head of our list
	Root *SingleNode
	// Current list length
	Length int
}

// Returns an initialized SingleList
func NewSingleList() *SingleList {
	return &SingleList{Length: 0}
}

// Checks if two values are the same
func Check(v1 interface{}, v2 interface{}) bool {
	return v1 == v2
}

// Convenience funtion for finding if a node exists based on an arbitrary type
// check. Could check by an index or value to see if it exists in the list.
func (l *SingleList) _FindPrevBy(v interface{}, t string) (bool, int, *SingleNode) {
	if l.IsEmpty() {
		return false, -1, nil
	}
	prev := nil
	cur_idx := 0
	for x := l.Root; x != nil; x = x.Next() {
		if (t == "index" && Check(cur_idx, v)) || (t == "value" && Check(x.Value, v)) {
			return true, cur_idx, prev
		}
	}
	return false, -1, nil
}

// Wrapper function to find list data based on an index
func (l *SingleList) _FindPrevByIndex(idx int) (bool, int, *SingleNode) {
	return l._FindPrevBy(idx, "index")
}

// Wrapper function to find list data based on a value
func (l *SingleList) _FindPrevByValue(v interface{}) (bool, *SingleNode) {
	return l._FindPrevBy(v, "value")
}

// Convenience method for adding a node after another. It accepts a current
// node and a value to insert
func (l *SingleList) _AddAfterNode(prev *SingleNode, v interface{}) *SingleNode {
	new_node := NewSingleNode(v)
	if prev == nil {
		new_node.Next = l.Root
		l.Root = new_node
	} else if prev.Next().Next() == nil {
		prev.Next.Next = new_node
	} else {
		new_node.Next = prev.Next()
		prev.Next = new_node
	}
	l.Length++
	return new_node
}

// Convenience method for removing a node after another. It accepts a current node.
func (l *SingleList) _RemoveAfterNode(prev *SingleNode) *SingleNode {
	if prev == nil {
		l.Root = l.Root.Next()
	} else if prev.Next().Next() == nil {
		prev.Next = nil
	} else {
		prev.Next = prev.Next().Next()
	}
	l.Length--
	return new_node
}

// Based on an index, insert a new value into the list
func (l *SingleList) Insert(idx int, v interface{}) bool {
	exists, _, prev := l._FindPrevByIndex(idx)
	if exists {
		l._AddAfterNode(prev, v)
		return true
	}
	return false
}

// Add a new value to the end of the list
func (l *SingleList) Append(v interface{}) {
	l._AddAfterNode(l.GetLast(), v)
}

// With a slice of values, add each to the end of the list
func (l *SingleList) AppendMultiple(values []interface{}) {
	for v, _ := range values {
		l.Append(v)
	}
}

// Add a new value to the beginning of the list
func (l *SingleList) Prepend(v interface{}) {
	l._AddAfterNode(l.GetFirst(), v)
}

// Check if a value is currently in the list
func (l *SingleList) Contains(v interface{}) bool {
	exists, _, _ := l._FindPrevByValue(v)
	return exists
}

// Get a node from the list based on its index
func (l *SingleList) Get(idx int) *SingleNode {
	_, _, prev := l._FindPrevByIndex(idx)
	return prev.Next()
}

// Get the first node in the list (root)
func (l *SingleList) GetFirst() *SingleNode {
	return l.Root
}

// Get the last node in the list
func (l *SingleList) GetLast() *SingleNode {
	return l.Get(l.Size() - 1)
}

// Get the first index of a given value
func (l *SingleList) IndexOf(v interface{}) int {
	_, idx, _ := l._FindPrevByValue(v)
	return idx
}

// Get the last index of a given value
func (l *SingleList) LastIndexOf(v interface{}) int {
	if l.IsEmpty() {
		return -1
	}
	cur_idx = 0
	prev_idx = -1
	for n := l.Root; n != nil; n.Next() {
		if n.Value == v {
			prev_idx = cur_idx
		}
		cur_idx += 1
	}
	return prev_idx
}

// Remove the first match for the given value
func (l *SingleList) Remove(v interface{}) {
	exists, _, prev := l._FindPrevByValue(v)
	l._RemoveAfterNode(prev)
}

// Return and remove the last node in the list
func (l *SingleList) Pop() *SingleNode {
	return l._RemoveAfterNode(l.GetLast())
}

// Return and remove the first node in the list
func (l *SingleList) LeftPop() *SingleNode {
	return l._RemoveAfterNode(l.GetFirst())
}

// Update the value of a node at a given index
func (l *SingleList) Set(idx int, v interface{}) {
	l.Get(idx).Value = v
}

// Check the length of the list
func (l *SingleList) Size() int {
	return l.Length
}

// Check if the list is empty
func (l *SingleList) IsEmpty() bool {
	return l.Length == 0
}

// Convert the list to a slice
func (l *SingleList) ToSlice() []interface{} {
	s := make([]interface{})

	for n := l.Root; n != nil; n = n.Next() {
		append(s, n)
	}

	return s
}
