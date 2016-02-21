package bst

// Node is an indiviual node in the BST which has a reference to its left and
// right subtrees and is placed in the BST based on it's value.
type Node struct {
	Val         int
	Left, Right *Node
}

// NewNode creates a new empty Node instance with both the left and right
// subtrees uninitialized.
func NewNode(val int) *Node {
	return &Node{Val: val}
}

// BST is a binary search tree beginning with a root node.
type BST struct {
	root *Node
	size int
}

// NewBST creates an empty instance of a binary search tree with.
func NewBST() *BST {
	return &BST{size: 0}
}

// Cardinality returns the current size of the BST.
func (bst *BST) Cardinality() int {
	return bst.size
}

// IsEmpty checks the cardinality of the BST for 0.
func (bst *BST) IsEmpty() bool {
	return bst.Cardinality() == 0
}

// Root returns the root node of a BST.
func (bst *BST) Root() *Node {
	return bst.root
}

// insert is an internal recursive function which does the work required to
// properly place a new node in a BST. It compares values of the new node to
// be inserted with the current node. If the new value is less than the
// current value and the left subtree is not empty, we recurse into that left
// subtree. Otherwise, we set the new node on that left subtree. The same goes
// for the right subtree for values greater than the current.
func (node *Node) insert(newNode *Node) bool {
	if newNode.Val < node.Val {
		if node.Left == nil {
			node.Left = newNode
			return true
		}
		return node.Left.insert(newNode)
	} else if newNode.Val > node.Val {
		if node.Right == nil {
			node.Right = newNode
			return true
		}
		return node.Right.insert(newNode)
	}
	return false
}

// Insert is the entry to the recursive insert function. If we have an empty
// BST, we set the root node. Otherwise, we recurse through the tree comparing
// nodes until it is placed. Unless the new value is a duplicate, it will be
// placed.
func (bst *BST) Insert(val int) bool {
	inserted := true
	if bst.IsEmpty() {
		bst.root = NewNode(val)
	} else {
		inserted = bst.Root().insert(NewNode(val))
	}
	if inserted {
		bst.size++
	}
	return inserted
}

// finMin is a supplementry function to deleting a node. When we are deleting
// a node with two children, we find the minimum value of the node's right
// subtree and return that an its parent in order to update the deleted node
// to this node.
func (node *Node) findMin(parent *Node) (*Node, *Node) {
	if node.Left == nil {
		return node, parent
	}
	return node.Left.findMin(node)
}

// link is a supplementry function to deleting a node. Based on the node being
// deleted, it sets the parents left or right subtree root as the new node.
// This allows us to properly remove a node and keep the child nodes under a
// deleted node.
func (node *Node) link(newNode, parent *Node) {
	if parent != nil {
		if node == parent.Left {
			parent.Left = newNode
		} else {
			parent.Right = newNode
		}
	}
	if newNode != nil {
		node = parent
	}
}

// del is the recursive supplementry function used when deleting a node. It
// checks to see if the values are less than, greater than, or equal to each
// other. If less than, we have to go deeper down the left subtree. If greater
// than, we go deeper down the right subtree. If they are equal, we determine
// whether the node has two children, one child, or no children. In the case
// of two children, we must find the minimum node in the right subtree and set
// that as the new parent to the child subtrees. We then must delete this
// value from its location in the child tree. If there is only one child, we
// accordingly link the parent node to this child to keep the connection. If
// it has no children, we can safely remove this node.
func (node *Node) del(parent *Node, val int) bool {
	switch {
	case node == nil:
		return false
	case val < node.Val:
		return node.Left.del(node, val)
	case val > node.Val:
		return node.Right.del(node, val)
	case val == node.Val:
		if node.Left != nil && node.Right != nil {
			// Handle node with two children
			next, nextParent := node.Right.findMin(node)
			node.Val = next.Val
			next.del(nextParent, next.Val)
		} else if node.Left != nil {
			// Handle node with only left child
			node.link(node.Left, parent)
		} else if node.Right != nil {
			// Handle node with only right child
			node.link(node.Right, parent)
		} else {
			// Handle node with no children
			node.link(nil, parent)
		}
		return true
	}
	return false
}

// Delete is the entry point to deleting a node on a BST. If the BST is empty,
// we don't have anything to delete. If it's the root node, we need to find
// the new root node. In any other case, we recurse down the left and right
// subtrees to see if the value exists to delete.
func (bst *BST) Delete(val int) bool {
	deleted := false
	if bst.IsEmpty() {
		return false
	}
	if bst.Root().Val == val {
		deleted = bst.Root().del(nil, val)
	} else {
		deleted = bst.Root().Left.del(bst.Root(), val) || bst.Root().Right.del(bst.Root(), val)
	}
	if deleted {
		bst.size--
	}
	return deleted
}

// Iter creates a channel used for iterating through a BST and using each
// value for the current node.
func (bst *BST) Iter(t Traversable) chan int {
	ch := make(chan int)
	go func() {
		t.Traverse(bst.Root(), ch)
		close(ch)
	}()
	return ch
}
