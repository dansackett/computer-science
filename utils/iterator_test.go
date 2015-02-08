package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	// It iterates
	s := []int{1, 2, 3, 4}
	it, _ := NewIterator(s)

	if it.Next() != 1 {
		fmt.Println("Wrong Value")
		t.Error()
	}

	it.Next()
	it.Next()
	it.Next()

	if it.Next() != nil {
		fmt.Println("Iteration didn't stop")
		t.Error()
	}

	// It only accepts slices
	_, err2 := NewIterator(6)
	if err2 == nil {
		fmt.Println(err2)
		t.Error()
	}

	// It sums stuff (just for fun)
	x := []int{1, 2, 3, 4}
	it3, _ := NewIterator(x)
	sum := 0
	for v := it3.Next(); v != nil; v = it3.Next() {
		fmt.Println(v)
		sum += v.(int)
	}

	if sum != 10 {
		fmt.Println(sum)
		t.Error()
	}
}
