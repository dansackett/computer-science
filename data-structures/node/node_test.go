package node

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	n := NewNode(10)

	// Test GetValue
	if n.GetValue() != 10 {
		fmt.Println(n.GetValue())
		t.Error()
	}

	// Test SetValue
	n.SetValue(20)
	if n.GetValue() != 20 {
		fmt.Println(n.GetValue())
		t.Error()
	}

	// Test SetNext and GetNext
	n2 := NewNode(30)
	n.SetNext(n2)
	if n.GetNext() != n2 {
		fmt.Println(n.GetNext())
		t.Error()
	}
}
