// ----------------------------------------------------------------------------
// Algorithms - Towers of Hanoi
// ----------------------------------------------------------------------------
//
// The Tower of Hanoi is a great problem in recursion involving three pegs and
// n number of blocks. The goal is to move the n blocks to the middle peg but
// there are rules:
//
// - You may only move one block at a time
// - A block may not have smaller blocks beneath it
// - A block can move to any peg
//
// Solving this with recursion is actually quite fun!
//
// ----------------------------------------------------------------------------
package recursion

import "fmt"

// Peg is one of the rods hoding a tower of blocks
type Peg struct {
	Blocks []int
}

// NewPeg inializes a new peg with some blocks
func NewPeg(blocks []int) *Peg {
	return &Peg{Blocks: blocks}
}

// Pop grabs the last value from a peg, removes it from the list, and returns it
func (p *Peg) Pop() int {
	val := p.Blocks[len(p.Blocks)-1]
	p.Blocks = p.Blocks[0 : len(p.Blocks)-1]
	return val
}

// Push places a new value onto the end of a peg
func (p *Peg) Push(val int) {
	p.Blocks = append(p.Blocks, val)
}

// MoveBlock pops a block off the source and pushes it onto the destination
func MoveBlock(source, dest *Peg) {
	dest.Push(source.Pop())
}

// Hanoi represents the set of pegs in the game
type Hanoi struct {
	Pegs map[string]*Peg
}

// NewHanoi initializes a new game of Hanoi by setting the number of blocks on
// peg `A` and none on pegs `B` and `C`
func NewHanoi(numBlocks int) *Hanoi {
	var initial []int
	for i := numBlocks; i > 0; i-- {
		initial = append(initial, i)
	}

	pegs := map[string]*Peg{
		"A": NewPeg(initial),
		"B": NewPeg([]int{}),
		"C": NewPeg([]int{}),
	}

	return &Hanoi{Pegs: pegs}
}

// GetPegs returns three pegs based on a `from` and `to` peg programatically
// determining the third peg based on those two
func (h *Hanoi) GetPegs(from, to string) (*Peg, *Peg, *Peg) {
	pegs := map[string]*Peg{
		"A": h.Pegs["A"],
		"B": h.Pegs["B"],
		"C": h.Pegs["C"],
	}

	fromPeg := pegs[from]
	delete(pegs, from)

	toPeg := pegs[to]
	delete(pegs, to)

	// Find the last peg in the set
	var key string
	for k := range pegs {
		key = k
	}
	otherPeg := pegs[key]

	return fromPeg, toPeg, otherPeg
}

// Print shows a basic display of the pegs in Hanoi in their current state
func (h *Hanoi) Print() {
	a, b, c := h.GetPegs("A", "B")
	pegs := []*Peg{a, b, c}

	fmt.Println("////////////////////")
	for index, peg := range pegs {
		fmt.Println("PEG", index+1)
		fmt.Println(peg.Blocks)
	}
}

// MoveForward is the recursion workhose which rotates the blocks around
func (h *Hanoi) MoveForward(index int, source, dest, spare *Peg) {
	if index == 0 {
		MoveBlock(source, dest)
	} else {
		h.MoveForward(index-1, source, spare, dest)
		MoveBlock(source, dest)
		h.MoveForward(index-1, spare, dest, source)
	}
}

// Run is the entry point of Hanoi allowing you to specify the number of
// blocks to use and the `from` and `to` pegs
func (h *Hanoi) Run(block int, from, to string) {
	fromPeg, toPeg, otherPeg := h.GetPegs(from, to)
	h.MoveForward(block-1, fromPeg, toPeg, otherPeg)
}
