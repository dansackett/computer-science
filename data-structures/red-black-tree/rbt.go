package rbt

import "fmt"

const (
	RED   = true
	BLACK = false
)

// Node is an indiviual node in the RBT which has a reference to its left and
// right subtrees and is placed in the RBT based on it's value.
type Node struct {
	Val                 int
	Color               bool
	Parent, Left, Right *Node
}

// LEAF is a sentinel node used to determine the default children for a node.
var LEAF = &Node{Color: BLACK, Val: -1}

// NewNode creates a new red Node instance with two leaves as children.
func NewNode(val int) *Node {
	return &Node{
		Val:   val,
		Color: RED,
		Left:  LEAF,
		Right: LEAF,
	}
}

// RBT is a red black tree beginning with a root node.
type RBT struct {
	root *Node
	size int
}

// NewRBT creates an empty instance of a red-black tree,
func NewRBT() *RBT {
	return &RBT{size: 0}
}

// IsRed reveals if the color of the passed in node is in fact Red.
func (node *Node) IsRed() bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}

// IsBlack reveals if the color of the passed in node is in fact Black.
func (node *Node) IsBlack() bool {
	return !node.IsRed()
}

// Cardinality returns the current size of the red-black tree.
func (rbt *RBT) Cardinality() int {
	return rbt.size
}

// IsEmpty checks the cardinality of the red-black tree.
func (rbt *RBT) IsEmpty() bool {
	return rbt.Cardinality() == 0
}

// Root returns the root node of a red-black tree.
func (rbt *RBT) Root() *Node {
	return rbt.root
}

// Find searches the red-black tree and returns the node if it finds it.
func (rbt *RBT) Find(val int) *Node {
	if rbt.IsEmpty() {
		return nil
	}
	node := rbt.Root()
	for node != LEAF {
		if node.Val == val {
			return node
		} else if node.Val < val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return nil
}

// String satisfies the println interface allowing us to print a
// representation of the tree. This is borrowed from:
// https://github.com/emirpasic/gods/blob/master/trees/redblacktree/redblacktree.go
func (rbt *RBT) String() string {
	str := "RedBlackTree\n"
	if !rbt.IsEmpty() {
		output(rbt.Root(), "", true, &str)
	}
	return str
}

// Insert attempts to insert the node into the red-black tree. It uses an
// internal recursive function to insert the node. Unless the node is a
// duplicate, it's expected to be inserted properly. Once inserted, the
// red-black tree rebalances itself to keep the red-black properties.
func (rbt *RBT) Insert(val int) bool {
	if inserted := rbt.insert(rbt.Root(), NewNode(val)); inserted != nil {
		rbt.size++
		rbt.insertBalance1(inserted)
		return true
	}
	return false
}

// Delete attempts to delete a node in the red-black tree. It first tries to
// find the node and if found, it checks how many children the node has. If it
// has two, we need to swap the lowest value in the right subtree and replace
// our node with it before attempting to balance the tree again. If the tree
// has zero or one child then it rebalances the tree.
func (rbt *RBT) Delete(val int) bool {
	if rbt.IsEmpty() {
		return false
	}

	node := rbt.Find(val)
	if node == nil {
		return false
	}

	// Handle node with two children
	if node.Left != LEAF && node.Right != LEAF {
		next := node.Right
		for next.Left != LEAF {
			next = next.Left
		}
		node.Val = next.Val
		node = next
	}

	// Find which child we're dealing with
	child := node.Right
	if node.Right == LEAF {
		child = node.Left
	}
	rbt.replaceNode(node, child)

	// Enter balancing cases
	if node.IsBlack() {
		if child.IsRed() {
			child.Color = BLACK
		} else {
			rbt.deleteBalance1(child)
		}
	}

	return true
}

// Iter creates a channel used for iterating through a red-black tree and
// using each value for the current node.
func (rbt *RBT) Iter(t Traversable) chan int {
	ch := make(chan int)
	go func() {
		t.Traverse(rbt.Root(), ch)
		close(ch)
	}()
	return ch
}

// replaceNode takes one node and places another node in its place.
func (rbt *RBT) replaceNode(curNode, newNode *Node) {
	newNode.Parent = curNode.Parent
	if curNode == curNode.Parent.Left {
		curNode.Parent.Left = newNode
	} else {
		curNode.Parent.Right = newNode
	}
}

// output is the recursive print tree function borrowed from:
// https://github.com/emirpasic/gods/blob/master/trees/redblacktree/redblacktree.go
func output(node *Node, prefix string, isTail bool, str *string) {
	if node.Right != LEAF {
		newPrefix := prefix
		if isTail {
			newPrefix += "|   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	color := "BLACK"
	if node.IsRed() {
		color = "RED"
	}
	*str += fmt.Sprintf("%v (%s)", node.Val, color) + "\n"
	if node.Left != LEAF {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "|   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

// insert is an internal recursive function which does the work required to
// properly place a new node in a tree. It compares values of the new node to
// be inserted with the current node. If the new value is less than the
// current value and the left subtree is not empty, we recurse into that left
// subtree. Otherwise, we set the new node on that left subtree. The same goes
// for the right subtree for values greater than the current.
func (rbt *RBT) insert(curNode, newNode *Node) *Node {
	if rbt.IsEmpty() {
		rbt.root = newNode
		return newNode
	} else if newNode.Val < curNode.Val {
		if curNode.Left == LEAF {
			curNode.Left = newNode
			newNode.Parent = curNode
			return newNode
		}
		return rbt.insert(curNode.Left, newNode)
	} else if newNode.Val > curNode.Val {
		if curNode.Right == LEAF {
			curNode.Right = newNode
			newNode.Parent = curNode
			return newNode
		}
		return rbt.insert(curNode.Right, newNode)
	}
	return nil
}

// grandparent represents the current node's parent's parent node. We need
// this in order to determine the insert and delete operations.
func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

// uncle represents the current node's "uncle" node which represents the
// sibling of the node's parent.
func (node *Node) uncle() *Node {
	g := node.grandparent()
	if g == nil {
		return nil
	}
	return node.Parent.sibling()
}

// sibling represents the current node's parent's sibling node,
func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

// insertBalance1 covers the balancing case 1 and is the recursive terminating
// case. If we reach the root of the tree then we paint it black again (if it
// had been changed) and end the loop. Otherwise, we go through the balancing
// cases again.
func (rbt *RBT) insertBalance1(node *Node) {
	if node.Parent == nil {
		node.Color = BLACK
	} else {
		rbt.insertBalance2(node)
	}
}

// insertBalance2 covers the balancing case 2 where a black parent node means
// we can safely insert the new child red node since the properties are
// unchanged. If the parent is red though, we must move to case 3.
func (rbt *RBT) insertBalance2(node *Node) {
	if node.Parent.IsBlack() {
		return
	}
	rbt.insertBalance3(node)
}

// insertBalance3 covers the balancing case 3 where we check on the uncle
// node. If the uncle is red then we can get away with simply flipping colors
// so the parent and uncle are now black and the grandparent is red. We
// recurse up to the grandparent node starting to check the balancing cases
// from there. If the uncle is black or these is no uncle though, we must move
// into case 4.
func (rbt *RBT) insertBalance3(node *Node) {
	u := node.uncle()

	if u != nil && u.IsRed() {
		node.Parent.Color = BLACK
		u.Color = BLACK
		g := node.grandparent()
		g.Color = RED
		rbt.insertBalance1(g)
	} else {
		rbt.insertBalance4(node)
	}
}

// insertBalance4 covers the balancing case 4 where we make our first
// rotations. If the current node is the parent's right and the parent is the
// grand parent's left then we must rotate it to the left. The opposite
// directions for the right apply as well for a right rotation.
func (rbt *RBT) insertBalance4(node *Node) {
	g := node.grandparent()

	if node == node.Parent.Right && node.Parent == g.Left {
		rbt.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == g.Right {
		rbt.rotateRight(node.Parent)
		node = node.Right
	}
	rbt.insertBalance5(node)
}

// insertBalance5 covers the final balancing case where we do a relabelling so
// we don't have two red nodes in a row from case 4. As well, it does a
// rotation around the grandparent node. This final rotation makes the parent
// node of our original node the new parent to the node and the grandparent.
// This case should have the tree / subtree fullfilling all properties of a
// red-black tree.
func (rbt *RBT) insertBalance5(node *Node) {
	g := node.grandparent()
	node.Parent.Color = BLACK
	g.Color = RED
	if node == node.Parent.Left {
		rbt.rotateRight(g)
	} else {
		rbt.rotateLeft(g)
	}
}

// deleteBalance1 covers the balancing case 1 which is also a terminating
// case. If we're at the root node, nothing happens. When we are in any
// subtree, we move into case 2.
func (rbt *RBT) deleteBalance1(node *Node) {
	if node.Parent != nil {
		rbt.deleteBalance2(node)
	}
}

// deleteBalance2 covers the balancing case 2 which checks if the sibling is a
// red node. If so then we can repaint the parent node red and the sibling
// node black followed by a rotation around the parent node. This is a setup
// case for other specialized cases.
func (rbt *RBT) deleteBalance2(node *Node) {
	sib := node.sibling()

	if sib.IsRed() {
		node.Parent.Color = RED
		sib.Color = BLACK
		if node == node.Parent.Left {
			rbt.rotateLeft(node.Parent)
		} else {
			rbt.rotateRight(node.Parent)
		}
	}
	rbt.deleteBalance3(node)
}

// deleteBalance3 covers the balancing case 3 which checks if the parent and
// the sibling subtree is black. If so, we simply paint the sibling red and go
// back to case 1 with the parent node. This fixes the property that a red
// node has two black children. If the subtree is more diverse, we move into
// case 4.
func (rbt *RBT) deleteBalance3(node *Node) {
	sib := node.sibling()

	if node.Parent.IsBlack() && sib.IsBlack() && sib.Left.IsBlack() && sib.Right.IsBlack() {
		sib.Color = RED
		rbt.deleteBalance1(node.Parent)
	} else {
		rbt.deleteBalance4(node)
	}
}

// deleteBalance4 covers the balancing case 4 which checks if the parent is
// red and the sibling subtree is black. This differs from case 3 because of
// the parent color. If this is the case, we only need to recolor the sibling
// as red and the parent as black to achieve what case 3 did. If the sibling
// tree and parent still do not match, we move into case 5.
func (rbt *RBT) deleteBalance4(node *Node) {
	sib := node.sibling()

	if node.Parent.IsRed() && sib.IsBlack() && sib.Left.IsBlack() && sib.Right.IsBlack() {
		sib.Color = RED
		node.Parent.Color = BLACK
	} else {
		rbt.deleteBalance5(node)
	}
}

// deleteBalance5 covers the balancing case 5 which checks if the sibling node
// is black. If so, it checks which subtree the current node is in relation to
// its parent. If it's the left subtree and the sibling's right node is black
// and left node is red then we can change the sibling node to red and it's
// left node to black and a right rotation will prepare us for case 6. If it's
// the right subtree and the sibling's left node is black and the right is red
// then we can do a left rotation setting the colors the same. In all cases,
// we move into case 6.
func (rbt *RBT) deleteBalance5(node *Node) {
	sib := node.sibling()

	if sib.IsBlack() {
		if node == node.Parent.Left && sib.Right.IsBlack() && sib.Left.IsRed() {
			sib.Color = RED
			sib.Left.Color = BLACK
			rbt.rotateRight(sib)
		} else if node == node.Parent.Right && sib.Left.IsBlack() && sib.Right.IsRed() {
			sib.Color = RED
			sib.Right.Color = BLACK
			rbt.rotateLeft(sib)
		}
	}

	rbt.deleteBalance6(node)
}

// deleteBalance6 covers the final balancing case which swaps the sibling's color
// with the parent's color and then sets the parent to black. Based on the
// relation of the node to the parent, we rotate left or right around the
// parent node to give us a final balanced state for the subtree.
func (rbt *RBT) deleteBalance6(node *Node) {
	sib := node.sibling()

	sib.Color = node.Parent.Color
	node.Parent.Color = BLACK

	if node == node.Parent.Left {
		sib.Right.Color = BLACK
		rbt.rotateLeft(node.Parent)
	} else {
		sib.Left.Color = BLACK
		rbt.rotateRight(node.Parent)
	}
}

// rotateLeft shifts the tree around a pivot node realigning the right subtree
// as the new root.
func (rbt *RBT) rotateLeft(node *Node) {
	y := node.Right
	node.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = node
	}
	y.Parent = node.Parent
	if node.Parent == nil {
		rbt.root = y
	} else if node == node.Parent.Left {
		node.Parent.Left = y
	} else {
		node.Parent.Right = y
	}
	y.Left = node
	node.Parent = y
}

// rotateRight shifts the tree around a pivot node realigning the left subtree
// as the new root.
func (rbt *RBT) rotateRight(node *Node) {
	y := node.Left
	node.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = node
	}
	y.Parent = node.Parent
	if node.Parent == nil {
		rbt.root = y
	} else if node == node.Parent.Right {
		node.Parent.Right = y
	} else {
		node.Parent.Left = y
	}
	y.Right = node
	node.Parent = y
}
