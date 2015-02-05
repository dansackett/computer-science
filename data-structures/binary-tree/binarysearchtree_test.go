package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	t1 := NewBST()

	// Test insert
	t1.Insert(16)
	t1.Insert(25)
	t1.Insert(42)

	if t1.Size() != 3 {
		fmt.Println(t1.Size())
		t.Error()
	}

	// Test clearkey
	t1.Root.ClearKey()
	if t1.Root.Key != -1 {
		fmt.Println(t1.Root.Key)
		t.Error()
	}

	// Test compare
	t1.Root.Key = 16
	c := Compare(t1.Root, t1.Root.Right)
	if c != -1 || c == 1 || c == 0 {
		fmt.Println(c)
		t.Error()
	}

	// Test lookup - Does exist
	n1 := t1.Lookup(25)
	if n1.Key != 25 {
		fmt.Println(n1.Key)
		t.Error()
	}

	// Test lookup - Does not exist
	n2 := t1.Lookup(4)
	if n2 != nil {
		fmt.Println(n2)
		t.Error()
	}

	// Test massinsert
	t2 := NewBST()
	x := []int{16, 25, 42, 8, 62, 49, 58}
	t2.MassInsert(x)
	if t2.Size() != 7 {
		fmt.Println(t2.Size())
		t.Error()
	}

	// Test remove
	t2.Remove(58)
	if found := t2.Lookup(58); found != nil {
		fmt.Println("8 exists")
		t.Error()
	}

	if t2.Size() != 6 {
		fmt.Println(t2.Size())
		t.Error()
	}

	// Test balance
	t2.Balance()
	if t2.Root.Key != 42 {
		fmt.Println(t2.Root.Key)
		t.Error()
	}
}
