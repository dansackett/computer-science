package bst

// Traversable is an interface to define a method for traversing through a BST.
type Traversable interface {
	Traverse(*Node, chan int)
}

// InOrderTraversable is an empty struct used to handle the methods of the
// Traversable interface.
type InOrderTraversable struct{}

// InOrder is a valid type used to specify the iteration method.
var InOrder InOrderTraversable

// Traverse satisfies the Traversable interface sending values back on the
// channel. This one if for in-order traversing.
func (bst InOrderTraversable) Traverse(node *Node, ch chan int) {
	if node != nil {
		bst.Traverse(node.Left, ch)
		ch <- node.Val
		bst.Traverse(node.Right, ch)
	}
}

// PreOrderTraversable is an empty struct used to handle the methods of the
// Traversable interface.
type PreOrderTraversable struct{}

// PreOrder is a valid type used to specify the iteration method.
var PreOrder PreOrderTraversable

// Traverse satisfies the Traversable interface sending values back on the
// channel. This one if for pre-order traversing.
func (bst PreOrderTraversable) Traverse(node *Node, ch chan int) {
	if node != nil {
		ch <- node.Val
		bst.Traverse(node.Left, ch)
		bst.Traverse(node.Right, ch)
	}
}

// PostOrderTraversable is an empty struct used to handle the methods of the
// Traversable interface.
type PostOrderTraversable struct{}

// PostOrder is a valid type used to specify the iteration method.
var PostOrder PostOrderTraversable

// Traverse satisfies the Traversable interface sending values back on the
// channel. This one if for post-order traversing.
func (bst PostOrderTraversable) Traverse(node *Node, ch chan int) {
	if node != nil {
		bst.Traverse(node.Left, ch)
		bst.Traverse(node.Right, ch)
		ch <- node.Val
	}
}
