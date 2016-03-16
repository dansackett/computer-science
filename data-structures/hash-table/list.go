package hash

type Node struct {
	key  string
	val  interface{}
	next *Node
}

type List struct {
	head, tail *Node
	size       int
}

func NewList() *List {
	return &List{size: 0}
}

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

func (l *List) Iter() chan *Node {
	ch := make(chan *Node)
	go func() {
		for cur := l.head; cur != nil; cur = cur.next {
			ch <- cur
		}
		close(ch)
	}()
	return ch
}
