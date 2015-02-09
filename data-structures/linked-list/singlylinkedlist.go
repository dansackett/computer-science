package singlylinkedlist

import (
	"errors"
	"github.com/dansackett/algorithms/data-structures/node"
)

type List struct {
	Length int
	Head   *node.Node
	Tail   *node.Node
}

func NewList() *List {
	return &List{Length: 0}
}

func (l *List) Append(v interface{}) {
	n := node.NewNode(v)
	if l.IsEmpty() {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	l.Length += 1
}

func (l *List) Prepend(v interface{}) {
	n := node.NewNode(v)
	if l.IsEmpty() {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
	l.Length += 1
}

func (l *List) AddAfter(n *node.Node, v interface{}) (bool, error) {
	if n == nil {
		return false, errors.New("Node does not exist")
	} else {
		m := node.NewNode(v)
		m.Next = n.Next
		n.Next = m
	}
	l.Length += 1
	return true, nil
}

func (l *List) Clear() {
	l.Head = nil
	l.Tail = nil
	l.Length = 0
}

func (l *List) Len() int {
	return l.Length
}

func (l *List) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List) Set(i int, v interface{}) (bool, error) {
	if l.IsEmpty() {
		return false, errors.New("List is empty")
	}
	if i >= l.Len() {
		return false, errors.New("Index out of range")
	}
	c := 0
	for n := l.Head; n != nil; n = n.Next {
		if c == i {
			n.Value = v
			return true, nil
		}
		c++
	}
	return false, errors.New("Index doesn't exist")
}

func (l *List) RemoveNode(n *node.Node) (bool, error) {
	if l.IsEmpty() {
		return false, errors.New("List is empty")
	}
	if n == nil {
		return false, errors.New("Node does not exist")
	}
	p := node.NewNode(0)
	for m := l.Head; m != nil; m = m.Next {
		if m == n {
			p.Next = m.Next
			l.Length -= 1
			return true, nil
		}
		p = m
	}
	return false, errors.New("Could not remove node")
}

func (l *List) RemoveIndex(i int) (bool, error) {
	if l.IsEmpty() {
		return false, errors.New("List is empty")
	}
	if i >= l.Len() {
		return false, errors.New("Index out of bounds")
	}
	c := 0
	p := node.NewNode(0)
	for n := l.Head; n != nil; n = n.Next {
		if c == i {
			p.Next = n.Next
			l.Length -= 1
			return true, nil
		}
		c++
	}
	return false, errors.New("Could not remove index")
}

func (l *List) Pop() (*node.Node, error) {
	if l.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	if l.Len() == 1 {
		h := l.Head
		l.Clear()
		return h, nil
	}
	p := node.NewNode(0)
	for n := l.Head; n != nil; n = n.Next {
		if n.Next == l.Tail {
			t := l.Tail
			n.Next = nil
			l.Tail = n
			l.Length -= 1
			return t, nil
		}
		p = n
	}
	p = nil
	return p, errors.New("Nothing to pop")
}

func (l *List) PopLeft() (*node.Node, error) {
	if l.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	if l.Len() == 1 {
		h := l.Head
		l.Clear()
		return h, nil
	}
	n := l.Head
	l.Head = n.Next
	l.Length -= 1
	return n, nil
}

func (l *List) Contains(v interface{}) (bool, error) {
	if l.IsEmpty() {
		return false, errors.New("List is empty")
	}
	for n := l.Head; n != nil; n = n.Next {
		if n.Value == v {
			return true, nil
		}
	}
	return false, errors.New("Item not in list")
}

func (l *List) Find(v interface{}) (*node.Node, error) {
	if l.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	for n := l.Head; n != nil; n = n.Next {
		if n.Value == v {
			return n, nil
		}
	}
	return nil, errors.New("Value not found")
}
