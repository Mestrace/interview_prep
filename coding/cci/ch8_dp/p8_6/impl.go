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

import (
	"sort"
)

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

func (h Hanoi) CanMove(i, j int) bool {
	// Make sure i and j are within index of towers
	if i >= len(h) || j >= len(h) {
		return false
	}
	// Condition:
	// 1. i-th tower must have elements
	// 2. the element to be moved in i-th tower must have a smaller value than
	//    the following element in j-th tower (or the j-th tower could be empty
	//    to hold value)
	if len(h[i]) != 0 && (len(h[j]) == 0 || h[i][len(h[i])-1] < h[j][len(h[j])-1]) {
		return true
	}

	return false
}

// MakeMove moves the top one from tower i to j; must call h.CanMove before make moves
func (h Hanoi) MakeMove(i, j int) {
	v := h[i][len(h[i])-1]
	h[i] = h[i][:len(h[i])-1]
	h[j] = append(h[j], v)
}

// HasWinned defines the winning state
func (h Hanoi) HasWinned() bool {
	for i := 0; i < len(h)-1; i++ {
		if len(h) != 0 {
			return false
		}
	}

	// is sorted on decreasing order?
	return sort.SliceIsSorted(h[len(h)-1], func(i, j int) bool {
		return h[len(h)-1][i] > h[len(h)-1][j]
	})
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

// func (m move) AllMoves(h Hanoi) []Hanoi {
// 	idxSet := make([]int, len(h))
// 	for i := range h {
// 		idxSet[i] = i
// 	}

// 	result := make([]Hanoi, 0, len(h))
// 	for _, comb := range Combination(idxSet) {
// 		result = append(result, move{
// 			State: h.Copy(),
// 			From: comb[0],
// 			From:
// 		})
// 	}
// }

func factorial(a int) int {
	if a <= 0 {
		panic("Negative number does no have factorial.")
	} else if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	}
	return a * factorial(a-1)
}

// Permutation returns all permutations of the given set
func Permutation(list []int) [][]int {
	var perm func(p []int, l, r int) [][]int
	perm = func(p []int, l, r int) [][]int {
		if l == r {
			return [][]int{p}
		}
		result := make([][]int, 0, r-l)
		for i := l; i <= r; i++ {
			p[l], p[i] = p[i], p[l]
			// captures p and return
			q := make([]int, len(p))
			copy(q, p)
			result = append(result, perm(q, l+1, r)...)
			p[l], p[i] = p[i], p[l]
		}
		return result
	}

	return perm(list, 0, len(list)-1)
}

// func SolveHanoiBFS(h Hanoi) Hanoi {
// 	stack := make([]move, 0)
// }
