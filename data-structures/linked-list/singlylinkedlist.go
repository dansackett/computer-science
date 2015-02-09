package singlylinkedlist

import (
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
	if l.Length == 0 {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

func (l *List) Prepend(v interface{}) {
	n := node.NewNode(v)
	if l.Length == 0 {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}

func (l *List) AddAfter(n *node.Node, v interface{}) {
	if n == nil {
		panic("Node does not exist!")
	} else {
		m := node.NewNode(v)
		m.Next = n.Next
		n.Next = m
	}
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

func (l *List) Set(i int, v interface{}) {
	c := 0
	for n := l.Head; n != nil; n = n.Next {
		if c == i {
			n.Value = v
			break
		}
		c++
	}
}

func (l *List) RemoveNode(n *node.Node) {
	p := node.NewNode(0)
	for m := l.Head; m != nil; m = m.Next {
		if m == n {
			p.Next = m.Next
			break
		}
		p = m
	}
}

func (l *List) RemoveIndex(i int) {
	c := 0
	p := node.NewNode(0)
	for n := l.Head; n != nil; n = n.Next {
		if c == i {
			p.Next = n.Next
			break
		}
		c++
	}
}

func (l *List) Pop() *node.Node {
	if l.Length == 0 {
		return nil
	}
	if l.Length == 1 {
		h := l.Head
		l.Clear()
		return h
	}
	p := node.NewNode(0)
	for n := l.Head; n != nil; n = n.Next {
		if n.Next == l.Tail {
			t := l.Tail
			n.Next = nil
			l.Tail = n
			return t
		}
		p = n
	}
	p = nil
	return p
}

func (l *List) PopLeft() *node.Node {
	if l.Length == 0 {
		return nil
	}
	if l.Length == 1 {
		h := l.Head
		l.Clear()
		return h
	}
	n := l.Head
	l.Head = n.Next
	return n
}

// func (l *List) Contains(v interface{}) bool {
// 	return true
// }

// func (l *List) Find(v interface{}) *node.Node {
// 	return
// }

// func (l *List) Iter() {
// }
