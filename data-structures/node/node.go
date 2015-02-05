/**
* Node
*
* A Node data structure is the simplest data structure in CS. The idea is it's
* an object that stores a value and a reference to the next Node. We can do
* the following operations with a Node:
*
* - Create Node
* - Get Value
* - Get Next
* - Set Value
* - Set Next
*
* All operations can be performed in O(1) time. This means all methods do not
* rely on any other variables.
**/
package node

import ()

type Node struct {
	Value int
	Next  *Node
}

func NewNode(v int) *Node {
	return &Node{Value: v}
}

func (n *Node) GetValue() int {
	return n.Value
}

func (n *Node) SetValue(v int) {
	n.Value = v
}

func (n *Node) GetNext() *Node {
	return n.Next
}

func (n *Node) SetNext(m *Node) {
	n.Next = m
}
