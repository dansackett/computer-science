package hash

// Node is a linked-list node
type Node struct {
	key  string
	val  interface{}
	next *Node
}

// List is a basic singly linked-list
type List struct {
	head, tail *Node
	size       int
}

// NewList returns an empty List
func NewList() *List {
	return &List{size: 0}
}

// GetHead returns the head of the list
func (l *List) GetHead() *Node {
	return l.head
}

// GetTail returns the tail of the list
func (l *List) GetTail() *Node {
	return l.tail
}

// GetHead returns the next value from the node
func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

// GetKey returns the node's key
func (n *Node) GetKey() string {
	return n.key
}

// GetVal returns the node's value
func (n *Node) GetVal() interface{} {
	return n.val
}

// Append adds a new value to the end of the list
func (l *List) Append(key string, val interface{}) {
	newNode := &Node{key: key, val: val}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// Remove removes a value based on its key from the list
func (l *List) Remove(key string) {
	if l.size == 0 {
		return
	}
	var cur, prev *Node
	for cur, prev = l.head, nil; cur != nil; cur, prev = cur.next, cur {
		if cur.key == key {
			if cur == l.head {
				l.head = l.head.next
			} else {
				prev.next = cur.next
			}
			l.size--
			return
		}
	}
}

// Iter returns a range-able iterator for easy traversal
func (l *List) Iter() chan *Node {
	ch := make(chan *Node)
	go func() {
		for cur := l.GetHead(); cur != nil; cur = cur.Next() {
			ch <- cur
		}
		close(ch)
	}()
	return ch
}
