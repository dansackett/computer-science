package recursion

import "testing"

func TestHanoiCanMoveFromFullTower(t *testing.T) {
	hanoi := NewHanoi(5)

	if len(hanoi.Pegs["A"].Blocks) != 5 {
		t.Errorf("`A` does not have the initial blocks")
	}

	hanoi.Run(5, "A", "B")

	if len(hanoi.Pegs["A"].Blocks) != 0 {
		t.Errorf("`A` has blocks that it should not have")
	}

	if len(hanoi.Pegs["B"].Blocks) != 5 {
		t.Errorf("`B` does not have the new blocks")
	}
}

func TestHanoiCanMovePartialTower(t *testing.T) {
	hanoi := NewHanoi(5)

	if len(hanoi.Pegs["A"].Blocks) != 5 {
		t.Errorf("`A` does not have the initial blocks")
	}

	hanoi.Run(3, "A", "B")

	if len(hanoi.Pegs["A"].Blocks) != 2 {
		t.Errorf("`A` has blocks that it should not have")
	}

	if len(hanoi.Pegs["B"].Blocks) != 3 {
		t.Errorf("`B` does not have the new blocks")
	}
}
