package linkedlist

import "testing"

func TestEmptyList(t *testing.T) {
	l := New()

	if l.Size() != 0 {
		t.Errorf("Initial list has length greater than 0, length is %d", l.Size())
	}

	if l.First() != l.Last() {
		t.Errorf("Initial list has head != tail")
	}
}

func TestListWithOneValue(t *testing.T) {
	l := New()
	l.Append(3)

	if l.Size() != 1 {
		t.Errorf("List should be of length 1, length is %d", l.Size())
	}

	if l.First().Value != 3 {
		t.Errorf("First value should be 3, first value is %d", l.First().Value)
	}

	if l.Last().Value != 3 {
		t.Errorf("Last value should be 3, last value is %d", l.Last().Value)
	}
}

func TestListWithMultipleValues(t *testing.T) {
	l := New()
	l.Append(3)
	l.Append(6)
	l.Prepend(9)

	// Test that we can set the first and last values of the list correctly
	if l.Size() != 3 {
		t.Errorf("List should be of length 3, length is %d", l.Size())
	}

	if l.First().Value != 9 {
		t.Errorf("First value should be 9, first value is %d", l.First().Value)
	}

	if l.Last().Value != 6 {
		t.Errorf("Last value should be 6, last value is %d", l.Last().Value)
	}

	// Test that we can remove nodes and the first and last values are OK
	l.Remove(l.First())
	l.Remove(l.Last())

	if l.Size() != 1 {
		t.Errorf("List should be of length 1, length is %d", l.Size())
	}

	if l.First().Value != 3 {
		t.Errorf("First value should be 3, first value is %d", l.First().Value)
	}

	if l.Last().Value != 3 {
		t.Errorf("Last value should be 3, last value is %d", l.Last().Value)
	}
}

func TestListCanBeAppendedWithList(t *testing.T) {
	l1 := New()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)

	l2 := New()
	l2.Append(4)
	l2.Append(5)
	l2.Append(6)

	l1.AppendList(l2)

	if l1.Size() != 6 {
		t.Errorf("List should be of length 6, length is %d", l1.Size())
	}

	if l1.First().Value != 1 {
		t.Errorf("First value should be 1, first value is %d", l1.First().Value)
	}

	if l1.Last().Value != 6 {
		t.Errorf("Last value should be 6, last value is %d", l1.Last().Value)
	}
}

func TestListCanBePrependedWithList(t *testing.T) {
	l1 := New()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)

	l2 := New()
	l2.Append(4)
	l2.Append(5)
	l2.Append(6)

	l1.PrependList(l2)

	if l1.Size() != 6 {
		t.Errorf("List should be of length 6, length is %d", l1.Size())
	}

	if l1.First().Value != 4 {
		t.Errorf("First value should be 4, first value is %d", l1.First().Value)
	}

	if l1.Last().Value != 3 {
		t.Errorf("Last value should be 3, last value is %d", l1.Last().Value)
	}
}

func TestListPopping(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	first := l.PopLeft()
	last := l.Pop()

	if l.Size() != 1 {
		t.Errorf("List should be of length 1, length is %d", l.Size())
	}

	if last.Value != 3 {
		t.Errorf("Popped value should be 3, value is %d", last.Value)
	}

	if first.Value != 1 {
		t.Errorf("Popped value should be 1, value is %d", first.Value)
	}

	if l.First() != l.Last() {
		t.Errorf("First node should now equal the last node")
	}
}

func TestListCanBeConvertedToSlice(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	s := l.ToSlice()
	vals := []int{1, 2, 3}
	target := make([]interface{}, len(vals))
	for i, x := range vals {
		target[i] = x
	}

	for i := range s {
		if s[i] != target[i] {
			t.Errorf("Slice values don't match, list is %d target is %d", s[i], target[i])
		}
	}
}
