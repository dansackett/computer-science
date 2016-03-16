package rbt

import "testing"

func TestInitialRBTIsEmpty(t *testing.T) {
	rbt := NewRBT()

	if !rbt.IsEmpty() {
		t.Error("RBT should be empty")
	}
}

func TestRBTWithRootHasTwoChildren(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)

	if rbt.IsEmpty() {
		t.Error("RBT should NOT be empty")
	}

	if rbt.Root().Left == nil || rbt.Root().Right == nil {
		t.Error("Node should always have two children")
	}
}

func TestRBTNodeCanDetermineColor(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)

	if rbt.Root().IsRed() {
		t.Error("Node should not be red")
	}

	if !rbt.Root().IsBlack() {
		t.Error("Node should be black")
	}
}

func TestRBTNodeCanGetParent(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)

	if rbt.Root().Right.Parent.Val != 1 {
		t.Errorf("Parent node should have a value of %d", 1)
	}
}

func TestRBTNodeCanGetGrandparent(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(3)
	rbt.Insert(4)

	if rbt.Find(4).grandparent().Val != 2 {
		t.Errorf("Grandparent node should have a value of %d", 2)
	}
}

func TestRBTNodeCanGetUncle(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(3)
	rbt.Insert(4)

	if rbt.Find(4).uncle().Val != 1 {
		t.Errorf("Uncle node should have a value of %d", 1)
	}
}

func TestRBTNodeCanGetSibling(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(3)
	rbt.Insert(4)

	if rbt.Find(3).sibling().Val != 1 {
		t.Errorf("Sibling node should have a value of %d", 1)
	}
}

func TestRBTCanGetRoot(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)

	if rbt.Root().Val != 1 {
		t.Errorf("Root node should have value of %d", 1)
	}
}

func TestRBTCanFindNodes(t *testing.T) {
	rbt := NewRBT()

	notFound := rbt.Find(10)

	if notFound != nil {
		t.Error("Should not find node in empty tree")
	}

	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(3)
	rbt.Insert(4)

	root := rbt.Find(2)

	if root.Val != 2 {
		t.Errorf("Found node should have value of %d, instead found %d", 2, root.Val)
	}

	left := rbt.Find(1)

	if left.Val != 1 {
		t.Errorf("Found node should have value of %d, instead found %d", 1, root.Val)
	}

	right := rbt.Find(4)

	if right.Val != 4 {
		t.Errorf("Found node should have value of %d, instead found %d", 4, root.Val)
	}
}

func TestRBTCanIterate(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(3)
	rbt.Insert(4)

	expected := []int{1, 2, 3, 4}

	i := 0
	for e := range rbt.Iter(InOrder) {
		if e != expected[i] {
			t.Errorf("Iteration should yield %d, found %d", expected[i], e)
		}
		i++
	}
}

func TestRBTCannotInsertDuplicateNode(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	res := rbt.Insert(1)

	if res {
		t.Error("Should not be allowed to insert duplicate node")
	}
}

func TestRBTInsertBalanceCase2(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)

	if rbt.Find(2).IsBlack() || rbt.Find(1).IsRed() {
		t.Error("Insert balance case 2 should not alter colors")
	}
}

func TestRBTInsertBalanceCase3(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(0)
	rbt.Insert(3)

	node := rbt.Find(3)

	if node.Parent.IsRed() {
		t.Error("Parent should be black after balance")
	}

	if node.uncle().IsRed() {
		t.Error("Uncle should be black after balance")
	}
}

func TestRBTInsertBalanceCase4(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(7)
	rbt.Insert(10)
	rbt.Insert(4)
	rbt.Insert(5)
	rbt.Insert(6)

	node := rbt.Find(6)

	if node.sibling().Val != 4 {
		t.Errorf("Sibling value should be 4, is %d", node.sibling().Val)
	}

	if node.Parent.Val != 5 {
		t.Errorf("Parent value should be 5, is %d", node.Parent.Val)
	}
}

func TestRBTCannotDeleteNodeNotInTree(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	res := rbt.Delete(2)

	if res {
		t.Error("Should not be allowed to delete node not in tree")
	}
}

func TestRBTCannotDeleteInEmptyTree(t *testing.T) {
	rbt := NewRBT()
	res := rbt.Delete(2)

	if res {
		t.Error("Should not be allowed to delete node in empty tree")
	}
}

func TestRBTCanDeleteNodeWithTwoChildren(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(3)
	res := rbt.Delete(2)

	if !res {
		t.Error("Should be allowed to delete node with two children")
	}

	if rbt.Root().Val != 3 {
		t.Errorf("New root node should be %d, found %d", 3, rbt.Root().Val)
	}

	if rbt.Root().IsRed() {
		t.Error("New root node should be black")
	}

	if rbt.Root().Right != LEAF {
		t.Error("New right node should be a LEAF")
	}

	if rbt.Root().Left.Val != 1 || !rbt.Root().Left.IsRed() {
		t.Error("New left node should be 1 and red")
	}
}

func TestRBTCanDeleteNodeWithOneChild(t *testing.T) {
	rbt := NewRBT()
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(3)
	rbt.Insert(4)
	rbt.Insert(5)
	rbt.Insert(6)
	res := rbt.Delete(5)

	if !res {
		t.Error("Should be allowed to delete node with one child")
	}
}

func TestRBTDeleteBalanceCase1(t *testing.T) {
}

func TestRBTDeleteBalanceCase2(t *testing.T) {
}

func TestRBTDeleteBalanceCase3(t *testing.T) {
}

func TestRBTDeleteBalanceCase4(t *testing.T) {
}

func TestRBTDeleteBalanceCase5(t *testing.T) {
}

func TestRBTDeleteBalanceCase6(t *testing.T) {
}
