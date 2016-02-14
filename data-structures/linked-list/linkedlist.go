package linkedlist

// Node represents a doubly linked node. As the spec states, we keep a
// reference to the next and previous nodes to avoid iteration.
type Node struct {
	next, prev *Node
	Value      interface{}
}

// Next returns the next node in the list unless there is no next node (it is
// the sentinel node) and then it returns nil
func (n *Node) Next() *Node {
	if e := n.next; e.Value != nil {
		return e
	}
	return nil
}

// Prev returns the previous node in the list unless there is no previous node
// (it is the sentinel node) and then it returns nil
func (n *Node) Prev() *Node {
	if e := n.prev; e.Value != nil {
		return e
	}
	return nil
}

// List represents a Circular Doubly Linked List.
type List struct {
	// sentinel is the value used to make this a circular list. It allows us
	// to keep a 0 index node in the list which points to the head and tail
	// easily.
	sentinel Node
	// size allows us to get the length of our list in O(1) time.
	size int
}

// Init is a way to initialize the list or clear it out if need be.
func (l *List) Init() *List {
	l.sentinel.next = &l.sentinel
	l.sentinel.prev = &l.sentinel
	l.size = 0
	return l
}

// New returns a new initialized List instance ready to be added to.
func New() *List {
	return new(List).Init()
}

// insertAfter is a convenience function which handles placing a node after a
// selected node. It also increases the list size and returns the new node
// instance.
func (l *List) insertAfter(newNode, node *Node) *Node {
	newNode.prev = node
	newNode.next = node.next
	node.next = newNode
	newNode.next.prev = newNode
	l.size++
	return newNode
}

// insertValue is a convenience function which calls insertAfter creating a
// new node in the process based on a passed in value.
func (l *List) insertValue(v interface{}, afterNode *Node) *Node {
	return l.insertAfter(&Node{Value: v}, afterNode)
}

// remove is a convenience function used to remove a specified node and
// decrease the list size.
func (l *List) remove(node *Node) *Node {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
	return node
}

// Size returns the length of the list.
func (l *List) Size() int {
	return l.size
}

// First returns the first node in the list unless the list is empty which it
// then returns nil.
func (l *List) First() *Node {
	if l.sentinel.next == &l.sentinel {
		return nil
	}
	return l.sentinel.next
}

// Last returns the last node in the list unless the list is empty which it
// then returns nil.
func (l *List) Last() *Node {
	if l.sentinel.prev == &l.sentinel {
		return nil
	}
	return l.sentinel.prev
}

// Append adds a new value to the end of the list.
func (l *List) Append(v interface{}) *Node {
	return l.insertValue(v, l.sentinel.prev)
}

// Prepend adds a new value to the beginnging of the list.
func (l *List) Prepend(v interface{}) *Node {
	return l.insertValue(v, &l.sentinel)
}

// Remove a node from the list if there are values in the list.
func (l *List) Remove(node *Node) interface{} {
	if l.size > 0 {
		l.remove(node)
	}
	return node.Value
}

// Pop removes the last value of the list returning the node.
func (l *List) Pop() *Node {
	if l.size == 0 {
		return nil
	}
	return l.remove(l.sentinel.prev)
}

// PopLeft removes the first value of the list returning the node.
func (l *List) PopLeft() *Node {
	if l.size == 0 {
		return nil
	}
	return l.remove(l.sentinel.next)
}

// AppendList adds another list to the end of the current list concatenating them.
func (l *List) AppendList(other *List) {
	for i, n := other.Size(), other.First(); i > 0; i, n = i-1, n.Next() {
		l.insertValue(n.Value, l.sentinel.prev)
	}
}

// PrependList adds another list to the beginning of the current list
// concatenating them.
func (l *List) PrependList(other *List) {
	for i, n := other.Size(), other.Last(); i > 0; i, n = i-1, n.Prev() {
		l.insertValue(n.Value, &l.sentinel)
	}
}

// ToSlice converts a List to a Golang slice
func (l *List) ToSlice() []interface{} {
	s := make([]interface{}, l.Size())

	for i, n := 0, l.First(); n != nil; i, n = i+1, n.Next() {
		s[i] = n.Value
	}

	return s
}
