package linkedlist

import ()

// SingleNode is a node in a SinglyLinkedList
type SingleNode struct {
	// Value stored for the node
	Value interface{}
	// Next item is a SingleNode which follows this one
	next *SingleNode
}

// Returns an initialized SingleNode
func NewSingleNode(v interface{}) *SingleNode {
	return &SingleNode{Value: v}
}

// Returns the item that follows a node
func (n *SingleNode) Next() *SingleNode {
	if n.next != nil {
		return n.next
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

// Function which takes variadic args and returns a new list of args to slice
func Slice(args ...interface{}) []interface{} {
	return args
}

// Convenience funtion for finding if a node exists based on an arbitrary type
// check. Could check by an index or value to see if it exists in the list.
func (l *SingleList) _FindPrevBy(v interface{}, t string) (bool, int, *SingleNode) {
	if l.IsEmpty() {
		return false, -1, nil
	}

	var prev *SingleNode
	prev = nil
	cur_idx := 0

	for x := l.GetFirst(); x != nil; x = x.Next() {
		if (t == "index" && Check(cur_idx, v)) || (t == "value" && Check(x.Value, v)) {
			return true, cur_idx, prev
		}
	}

	return false, -1, nil
}

// Wrapper function to find list data based on an index
func (l *SingleList) _FindPrevByIndex(idx int) (bool, int, *SingleNode) {
	if idx > l.Size()-1 {
		return false, -1, nil
	}

	return l._FindPrevBy(idx, "index")
}

// Wrapper function to find list data based on a value
func (l *SingleList) _FindPrevByValue(v interface{}) (bool, int, *SingleNode) {
	return l._FindPrevBy(v, "value")
}

// Convenience method for adding a node after another. It accepts a current
// node and a value to insert
func (l *SingleList) _AddAfterNode(prev *SingleNode, v interface{}) *SingleNode {
	new_node := NewSingleNode(v)

	if prev == nil {
		new_node.next = l.GetFirst()
		l.Root = new_node
	} else if prev.Next().Next() == nil {
		prev.next.next = new_node
	} else {
		new_node.next = prev.Next()
		prev.next = new_node
	}

	l.Length++

	return new_node
}

// Convenience method for removing a node after another. It accepts a current node.
func (l *SingleList) _RemoveAfterNode(prev *SingleNode) *SingleNode {
	var old_node *SingleNode

	if prev == nil {
		old_node = l.GetFirst()
		l.Root = l.GetFirst().Next()
	} else if prev.Next().Next() == nil {
		old_node = l.GetLast()
		prev.next = nil
	} else {
		old_node = prev.Next()
		prev.next = prev.Next().Next()
	}

	l.Length--

	return old_node
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
func (l *SingleList) Append(v interface{}) *SingleNode {
	return l._AddAfterNode(l.GetLast(), v)
}

// With a slice of values, add each to the end of the list
func (l *SingleList) AppendMultiple(values []interface{}) {
	for v, _ := range values {
		l.Append(v)
	}
}

// Add a new value to the beginning of the list
func (l *SingleList) Prepend(v interface{}) *SingleNode {
	return l._AddAfterNode(l.GetFirst(), v)
}

// Check if a value is currently in the list
func (l *SingleList) Contains(v interface{}) bool {
	return Slice(l._FindPrevByValue(v))[0].(bool)
}

// Get a node from the list based on its index
func (l *SingleList) Get(idx int) *SingleNode {
	n := Slice(l._FindPrevByIndex(idx))[2].(*SingleNode)

	if n != nil {
		return n.Next()
	}

	return l.GetFirst()
}

// Get the first node in the list (root)
func (l *SingleList) GetFirst() *SingleNode {
	return l.Root
}

// Get the last node in the list
func (l *SingleList) GetLast() *SingleNode {
	if l.IsEmpty() {
		return nil
	}
	return l.Get(l.Size() - 1)
}

// Get the first index of a given value
func (l *SingleList) IndexOf(v interface{}) int {
	return Slice(l._FindPrevByValue(v))[1].(int)
}

// Get the last index of a given value
func (l *SingleList) LastIndexOf(v interface{}) int {
	if l.IsEmpty() {
		return -1
	}

	cur_idx := 0
	prev_idx := -1

	for n := l.GetFirst(); n != nil; n.Next() {
		if Check(n.Value, v) {
			prev_idx = cur_idx
		}
		cur_idx++
	}

	return prev_idx
}

// Remove the first match for the given value
func (l *SingleList) Remove(v interface{}) *SingleNode {
	return l._RemoveAfterNode(Slice(l._FindPrevByValue(v))[2].(*SingleNode))
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
	return l.Size() == 0
}

// Convert the list to a slice
func (l *SingleList) ToSlice() []interface{} {
	s := make([]interface{}, l.Size())

	ctr := 0
	for n := l.GetFirst(); n != nil; n = n.Next() {
		s[ctr] = n
		ctr++
	}

	return s
}
