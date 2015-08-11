/**
* Binary Search Tree
*
* A Binary Search Tree is a type of Binary Tree which follows the rules:
*
* - Left children are smaller than the root node
* - Right children are larger than the root node
*
* These are great for searching as they reduce the amount of items to check.
* Still, they can be challenging to implement as I've learned. The following
* is my attempt at creating a self-balancing binary search tree. For this to
* be true, the following must also be true:
*
* - Upon Insertion and Deletion, the BST reshuffles itself so not to become	unbalanced.
* - At all times, the BST follows standard BST principles laid out above.
*
* To do this, we have a BST struct and a Node struct. a BST is a collection of
* nodes with a specified root. We can do the following to our BST:
*
* - Find a Node by key
* - Insert a new Node
* - Delete a Node
* - Balance the BST
* - Get size of BST
*
* Time Complexity:
*
*			Average		Worst case
* Space		O(n)		O(n)
* Search	O(log n)	O(n)
* Insert	O(log n)	O(n)
* Delete	O(log n)	O(n)
**/

package binarysearchtree

import (
	"math"
)

type Node struct {
	Key    int
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(k int) *Node {
	return &Node{Key: k}
}

func (n *Node) ClearKey() {
	n.Key = -1
}

func Compare(n *Node, m *Node) int {
	if n.Key < m.Key {
		return -1
	} else if n.Key > m.Key {
		return 1
	} else {
		return 0
	}
}

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{}
}

func (t *BST) _Insert(newNode *Node, root *Node) {
	// Set root if doesn't exist
	if t.Root == nil {
		t.Root = newNode
		return
	}

	c := Compare(newNode, root)

	if c <= 0 {
		// If new Node key < root Node key, then it belongs in left subtree
		if root.Left == nil {
			root.Left = newNode
			newNode.Parent = root
		} else {
			t._Insert(newNode, root.Left)
		}
	} else if c >= 0 {
		// If new Node key > root Node key, then it belongs in right subtree
		if root.Right == nil {
			root.Right = newNode
			newNode.Parent = root
		} else {
			t._Insert(newNode, root.Right)
		}
	}
}

func (t *BST) Insert(k int) {
	n := NewNode(k)
	t._Insert(n, t.Root)
}

func (t *BST) MassInsert(s []int) {
	for _, v := range s {
		t.Insert(v)
	}
}

func (t *BST) _Lookup(k int, r *Node) *Node {
	if r == nil {
		return nil
	}

	n := NewNode(k)
	c := Compare(n, r)

	// If we have a match, return the matching Node
	if c == 0 {
		return r
	}

	if c < 0 {
		return t._Lookup(k, r.Left)
	} else if c > 0 {
		return t._Lookup(k, r.Right)
	}

	return nil
}

func (t *BST) Lookup(k int) *Node {
	if t.Root == nil {
		return nil
	}
	return t._Lookup(k, t.Root)
}

func (t *BST) Contains(k int) bool {
	f := t.Lookup(k)
	if f != nil {
		return true
	}
	return false
}

func (t *BST) Remove(k int) bool {
	if t.Root == nil {
		return false
	}

	n := t.Lookup(k)

	if n == nil {
		return false
	}

	n.ClearKey()

	return true
}

func (t *BST) CreateSlice(n *Node) []int {
	s := []int{}

	if n == nil {
		return s
	}

	left := t.CreateSlice(n.Left)
	right := t.CreateSlice(n.Right)

	if n.Key == -1 {
		s = append(s, left...)
		s = append(s, right...)
		return s
	}

	// This is an InOrder representation
	s = append(s, left...)
	s = append(s, n.Key)
	s = append(s, right...)

	return s
}

func (t *BST) List() []int {
	return t.CreateSlice(t.Root)
}

func (t *BST) _Balance(s []int) {
	if len(s) == 0 {
		return
	}

	// Find middle
	mi := int32(math.Ceil(float64(len(s))/2.0)) - 1
	m := s[mi]
	s = append(s[:mi], s[mi+1:]...)
	c1 := s[:mi]
	c2 := s[mi:]

	t.Insert(m)
	t._Balance(c1)
	t._Balance(c2)
}

func (t *BST) Balance() {
	var m int
	s := t.CreateSlice(t.Root)

	if len(s) == 1 {
		t.Insert(s[0])
	} else {
		// Find middle
		mi := int32(math.Ceil(float64(len(s))/2.0)) - 1
		t.Root = nil

		if len(s)%2 == 0 {
			m = s[mi+1]
			s = append(s[:mi], s[mi+1:]...)
		} else {
			m = s[mi]
			s = append(s[:mi], s[mi+1:]...)
		}

		// Split list into chunks
		c1 := s[:mi]
		c2 := s[mi:]

		t.Insert(m)
		t._Balance(c1)
		t._Balance(c2)
	}
}

func (t *BST) Size() int {
	s := t.CreateSlice(t.Root)
	return len(s)
}
