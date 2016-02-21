package bst

import (
	"fmt"
	"strings"
)

// debugPrint is a test function use to visualize depth when debugging
// procedures within the tree.
func debugPrint(n *Node, ord string, depth int) {
	if n != nil {
		fmt.Printf("%s %s - %d\n", strings.Repeat("-", depth*3), ord, n.Val)
		debugPrint(n.Left, "LEFT", depth+1)
		debugPrint(n.Right, "RIGHT", depth+1)
	}
}
