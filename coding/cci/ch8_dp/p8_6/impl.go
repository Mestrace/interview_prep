/*
Package p8_6 solves the following problem:
In the classic Problem of the Twoers of Hanoi, you have 3 towers and N disks of
different sizes which can slide onto any tower. The puzzle starts with disks sorted
in ascending order of size from top to bottom (i.e, ech disk sits on top of an
even larger one). You have the following constraints:
  1. Only one disk can be moved at a time
  2. A disk is slid off the top of one tower onto another tower
  3. A disk cannot be placed on top of a smaller disk.
Write a program to move the disks from the first tower to the last using stacks.
*/
package p8_6

import "sort"

// Hanoi is a Towers of Hanoi game
type Hanoi [][]int

// Valid returns true when the Hanoi game is valid
func (h Hanoi) Valid() bool {
	for _, t := range h {
		if !sort.SliceIsSorted(t, func(i, j int) bool {
			return t[i] < t[j]
		}) {
			return false
		}
	}
	return true
}

// Move moves the top one from tower i to j, and returns true if the move
// can and was made
func (h Hanoi) Move(i, j int) bool {
	if i >= len(h) || j >= len(h) {
		return false
	}

	if len(h[i]) != 0 && (len(h[j]) == 0 || h[i][len(h[i])-1] < h[j][len(h[j])-1]) {
		v := h[i][len(h[i])-1]
		h[i] = h[i][:len(h[i])-1]
		h[j] = append(h[j], v)
		return true
	}
	return false
}

// HasWinned defines the winning state
func (h Hanoi) HasWinned() bool {
	for i := 0; i < len(h)-1; i++ {
		if len(h) != 0 {
			return false
		}
	}

	return sort.IntsAreSorted(h[len(h)-1])
}

// Copy returns a copy of current states
func (h Hanoi) Copy() Hanoi {
	result := make(Hanoi, len(h))

	for i := range h {
		result[i] = make([]int, len(h[i]))
		copy(result[i], h[i])
	}

	return result
}

type move struct {
	State Hanoi
	From  int
	To    int
}

func (m move) AllMoves(h Hanoi) {
	
}

func SolveHanoiBFS(h Hanoi) Hanoi {
	stack := make([]move, 0)
}
