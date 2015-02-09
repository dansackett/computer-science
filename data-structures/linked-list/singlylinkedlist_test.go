package singlylinkedlist

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	// Can create
	l1 := NewList()
	if l1.Length != 0 {
		fmt.Println("List initialized incorrectly")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can append
	l1.Append(20)
	if l1.Length != 1 && l1.Head.Value != 20 && l1.Tail.Value != 20 {
		fmt.Println(l1)
		fmt.Println("Node not appended")
		t.Error()
	}

	// ------------------------------------------------------------------------

	l1.Append(23)
	if l1.Length != 2 && l1.Head.Value != 20 && l1.Tail.Value != 23 {
		fmt.Println(l1)
		fmt.Println("Node not appended")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can prepend
	l1.Prepend(2)
	if l1.Length != 3 && l1.Head.Value != 2 && l1.Tail.Value != 23 {
		fmt.Println(l1)
		fmt.Println("Node not prepended")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can addafter
	l1.AddAfter(l1.Head, 5)
	if l1.Length != 4 && l1.Head.Next.Value != 5 {
		fmt.Println(l1)
		fmt.Println("Node not added after head")
		t.Error()
	}

	// Cannot add after
	b, _ := l1.AddAfter(nil, 5)
	if l1.Length != 4 && b != false {
		fmt.Println(l1)
		fmt.Println("Node should not have been added after")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can clear
	l1.Clear()
	if l1.Length != 0 && l1.Head != nil && l1.Tail != nil {
		fmt.Println(l1)
		fmt.Println("List not cleared")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can get length
	if l1.Len() != 0 {
		fmt.Println(l1)
		fmt.Println("Cannot get length")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can check empty
	if l1.IsEmpty() != true {
		fmt.Println(l1)
		fmt.Println("Cannot check if empty")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Cannot set index for empty list
	b2, _ := l1.Set(2, 7)
	if b2 == false && l1.Len() != 0 {
		fmt.Println(l1)
		fmt.Println("Set index, but doesn't exist")
		t.Error()
	}

	// Can set index
	l1.Append(3)
	l1.Append(4)
	l1.Append(5)
	b3, _ := l1.Set(2, 7)
	if b3 == true && l1.Head.Next.Value == 2 && l1.Len() != 3 {
		fmt.Println(l1)
		fmt.Println("Cannot set index")
		t.Error()
	}

	// Cannot set index out of range
	l1.Append(3)
	l1.Append(4)
	l1.Append(5)
	b4, _ := l1.Set(10, 7)
	if b4 == false && l1.Len() != 6 {
		fmt.Println(l1)
		fmt.Println("Can set index out of range")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Cannot remove node in empty list
	l1.Clear()
	b5, _ := l1.RemoveNode(l1.Head)
	if b5 == false && l1.Len() != 0 {
		fmt.Println(l1)
		fmt.Println("Removed node in empty list")
		t.Error()
	}

	// Can remove node
	l1.Append(3)
	l1.Append(4)
	l1.Append(5)
	b6, _ := l1.RemoveNode(l1.Head)
	if b6 == true && l1.Head.Value == 4 && l1.Len() != 2 {
		fmt.Println(l1)
		fmt.Println("Cannot remove head node")
		t.Error()
	}

	// Cannot remove non-existant node
	b7, _ := l1.RemoveNode(nil)
	if b7 == false && l1.Len() != 2 {
		fmt.Println(l1)
		fmt.Println("Can remove non-existant node")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Cannot remove index in empty list
	l1.Clear()
	b8, _ := l1.RemoveIndex(4)
	if b8 == false && l1.Len() != 0 {
		fmt.Println(l1)
		fmt.Println("Removed index in empty list")
		t.Error()
	}

	// Can remove index
	l1.Append(3)
	l1.Append(4)
	l1.Append(5)
	b9, _ := l1.RemoveIndex(2)
	if b9 == true && l1.Head.Next.Value == 5 && l1.Len() != 2 {
		fmt.Println(l1)
		fmt.Println("Cannot remove index")
		t.Error()
	}

	// Cannot remove non-existant index
	b10, _ := l1.RemoveIndex(10)
	if b10 == false && l1.Len() != 2 {
		fmt.Println(l1)
		fmt.Println("Can remove non-existant index")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can pop node
	n, _ := l1.Pop()
	if n.Value == 5 && l1.Len() != 1 {
		fmt.Println(l1)
		fmt.Println("Cannot pop")
		t.Error()
	}

	// Cannot pop from empty list
	l1.Clear()
	n2, _ := l1.Pop()
	if n2 == nil && l1.Len() != 0 {
		fmt.Println(l1)
		fmt.Println("Can pop from empty list")
		t.Error()
	}

	// Pops head for list of length 1
	l1.Append(10)
	n3, _ := l1.Pop()
	if n3.Value == 10 && l1.Len() != 0 {
		fmt.Println(l1)
		fmt.Println("Pop on list length 1 != Head")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can popleft node
	l1.Clear()
	l1.Append(4)
	l1.Append(5)
	l1.Append(6)
	n4, _ := l1.PopLeft()
	if n4.Value == 4 && l1.Len() != 2 {
		fmt.Println(l1)
		fmt.Println("Cannot popleft")
		t.Error()
	}

	// Cannot popleft from empty list
	l1.Clear()
	n5, _ := l1.PopLeft()
	if n5 == nil && l1.Len() != 0 {
		fmt.Println(l1)
		fmt.Println("Can popleft from empty list")
		t.Error()
	}

	// Popleft head for list of length 1
	l1.Append(10)
	n6, _ := l1.PopLeft()
	if n6.Value == 10 && l1.Len() != 0 {
		fmt.Println(l1)
		fmt.Println("Popleft on list length 1 != Head")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Empty list doesn't contain anything
	l1.Clear()
	b12, _ := l1.Contains(10)
	if b12 != false {
		fmt.Println(l1)
		fmt.Println("Contains in empty list")
		t.Error()
	}

	// Contains number
	l1.Append(10)
	l1.Append(11)
	l1.Append(12)
	l1.Append(13)
	b13, _ := l1.Contains(12)
	if b13 != true {
		fmt.Println(l1)
		fmt.Println("Contains doesn't work")
		t.Error()
	}

	// Does not contain number
	l1.Append(10)
	l1.Append(11)
	l1.Append(12)
	l1.Append(13)
	b14, _ := l1.Contains(15)
	if b14 != false {
		fmt.Println(l1)
		fmt.Println("Contains doesn't work")
		t.Error()
	}

	// ------------------------------------------------------------------------

	// Can't find in empty list
	l1.Clear()
	n7, _ := l1.Find(10)
	if n7 != nil {
		fmt.Println(l1)
		fmt.Println("Found in empty list")
		t.Error()
	}

	// Finds in list
	l1.Append(10)
	l1.Append(11)
	l1.Append(12)
	l1.Append(13)
	n8, _ := l1.Find(12)
	if n8.Value != 12 && l1.Len() != 4 {
		fmt.Println(l1)
		fmt.Println("Find doesn't work")
		t.Error()
	}

	// Finds in list
	n9, _ := l1.Find(15)
	if n9 != nil && l1.Len() != 4 {
		fmt.Println(l1)
		fmt.Println("Find doesn't work")
		t.Error()
	}
}
