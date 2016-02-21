package bst

import "testing"

func TestEmptyBST(t *testing.T) {
	bst := NewBST()

	if !bst.IsEmpty() {
		t.Errorf("BST should be empty")
	}
}

func TestSingleNodeBST(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)

	if bst.IsEmpty() {
		t.Errorf("BST should be NOT empty")
	}

	if bst.Cardinality() != 1 {
		t.Errorf("BST should have 1 item, has %d", bst.Cardinality())
	}

	if bst.Root().Val != 1 {
		t.Errorf("Root node value should be equal to 1, is %d", bst.Root().Val)
	}
}

func TestBSTSubtrees(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(20)

	if bst.Cardinality() != 3 {
		t.Errorf("BST should have 3 items, has %d", bst.Cardinality())
	}

	if bst.Root().Val != 10 {
		t.Errorf("Root node value should be equal to 10, is %d", bst.Root().Val)
	}

	if bst.Root().Left.Val != 5 {
		t.Errorf("Root node value should be equal to 5, is %d", bst.Root().Left.Val)
	}

	if bst.Root().Right.Val != 20 {
		t.Errorf("Root node value should be equal to 20, is %d", bst.Root().Right.Val)
	}
}

func TestBSTDoesNotAllowDuplicates(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	res := bst.Insert(10)

	if res {
		t.Errorf("Added duplicate but should not have")
	}

	if bst.Cardinality() != 1 {
		t.Errorf("BST should have 1 item, has %d", bst.Cardinality())
	}
}

func TestEmptyBSTCannotDelete(t *testing.T) {
	bst := NewBST()
	res := bst.Delete(1)

	if res {
		t.Errorf("Deleted from an empty BST")
	}
}

func TestBSTCanDeleteRoot(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(0)
	res := bst.Delete(1)

	if !res {
		t.Errorf("Tried to delete root but could not")
	}

	if bst.Cardinality() != 2 {
		t.Errorf("BST should have 2 items, has %d", bst.Cardinality())
	}

	if bst.Root().Val != 2 {
		t.Errorf("New root node should be 2, is %d", bst.Root().Val)
	}

	if bst.Root().Left.Val != 0 {
		t.Errorf("New root node's left node should be 0, is %d", bst.Root().Left.Val)
	}

	if bst.Root().Right != nil {
		t.Errorf("New root node's right node should nil")
	}
}

func TestBSTCanDeleteParentWithTwoChildren(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(0)
	bst.Insert(2)
	bst.Insert(6)
	res := bst.Delete(5)

	if !res {
		t.Errorf("Tried to delete 5 but could not")
	}

	if bst.Cardinality() != 4 {
		t.Errorf("BST should have 4 items, has %d", bst.Cardinality())
	}

	if bst.Root().Right.Val != 6 {
		t.Errorf("New node should be 6, is %d", bst.Root().Right.Val)
	}

	if bst.Root().Right.Left.Val != 2 {
		t.Errorf("New node's left node should be 2, is %d", bst.Root().Right.Left.Val)
	}

	if bst.Root().Right.Right != nil {
		t.Errorf("New node's right node should be nil")
	}
}

func TestBSTCanDeleteParentWithOneChild(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(0)
	bst.Insert(2)
	res := bst.Delete(5)

	if !res {
		t.Errorf("Tried to delete 5 but could not")
	}

	if bst.Cardinality() != 3 {
		t.Errorf("BST should have 3 items, has %d", bst.Cardinality())
	}

	if bst.Root().Right.Val != 2 {
		t.Errorf("New node should be 2, is %d", bst.Root().Right.Val)
	}

	if bst.Root().Right.Left != nil {
		t.Errorf("New node's left node should be nil")
	}

	if bst.Root().Right.Right != nil {
		t.Errorf("New node's right node should be nil")
	}
}

func TestBSTCanDeleteParentWithNoChildren(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(0)
	res := bst.Delete(0)

	if !res {
		t.Errorf("Tried to delete 0 but could not")
	}

	if bst.Cardinality() != 2 {
		t.Errorf("BST should have 2 items, has %d", bst.Cardinality())
	}

	if bst.Root().Left != nil {
		t.Errorf("New node's left node should be nil")
	}

	if bst.Root().Right == nil {
		t.Errorf("New node's right node should not be nil")
	}
}

func TestBSTCanIterInOrder(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(0)

	expected := []int{0, 1, 5}

	idx := 0
	for e := range bst.Iter(InOrder) {
		if e != expected[idx] {
			t.Errorf("InOrder iteration returned %d should have been %d", e, expected[idx])
		}
		idx++
	}
}

func TestBSTCanIterPreOrder(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(0)

	expected := []int{1, 0, 5}

	idx := 0
	for e := range bst.Iter(PreOrder) {
		if e != expected[idx] {
			t.Errorf("PreOrder iteration returned %d should have been %d", e, expected[idx])
		}
		idx++
	}
}

func TestBSTCanIterPostOrder(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(0)

	expected := []int{0, 5, 1}

	idx := 0
	for e := range bst.Iter(PostOrder) {
		if e != expected[idx] {
			t.Errorf("PostOrder iteration returned %d should have been %d", e, expected[idx])
		}
		idx++
	}
}
