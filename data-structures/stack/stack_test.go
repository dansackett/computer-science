package stack

import "testing"

func TestInitialStackIsEmpty(t *testing.T) {
	stack := NewStack()

	if !stack.IsEmpty() {
		t.Errorf("Stack is supposed to be empty")
	}
}

func TestStackWorks(t *testing.T) {
	stack := NewStack()
	stack.Push(5)
	stack.Push(8)
	stack.Push(1)

	if stack.Size() != 3 {
		t.Errorf("Stack should have 3 items in it, it has %d", stack.Size())
	}

	first := stack.Pop()

	if stack.Size() != 2 {
		t.Errorf("Stack should have 2 items in it, it has %d", stack.Size())
	}

	if first != 1 {
		t.Errorf("First value should be 1, it returned %d", first)
	}

	second := stack.Pop()
	third := stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("Stack is supposed to be empty")
	}

	if second != 8 {
		t.Errorf("Second value should be 8, it returned %d", second)
	}

	if third != 5 {
		t.Errorf("Third value should be 5, it returned %d", third)
	}
}
