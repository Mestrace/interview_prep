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
	"bytes"
	"encoding/gob"
	"fmt"
	"hash"
	"sort"

	"github.com/twmb/murmur3"
)

var (
	hasher hash.Hash64
)

func init() {
	hasher = murmur3.New64()
}

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

type Move struct {
	State       Hanoi
	From        int
	To          int
	ParentState *Move
}

func (m Move) CanMove() bool {
	return m.State.CanMove(m.From, m.To)
}

func (m Move) MakeMove() {
	m.State.MakeMove(m.From, m.To)
}

func (m Move) HasWinned() bool {
	return m.State.HasWinned()
}

func (m Move) AllMoves() []Move {
	idxSet := make([]int, len(m.State))
	for i := range m.State {
		idxSet[i] = i
	}

	result := make([]Move, 0, len(m.State))
	for _, comb := range Combination(idxSet, 2) {
		result = append(result, Move{
			State:       m.State.Copy(),
			From:        comb[0],
			To:          comb[1],
			ParentState: &m,
		})
	}
	return result
}

func (m Move) Hash() uint64 {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	e.Encode(m.State)
	e.Encode(m.From)
	e.Encode(m.To)

	hasher.Reset()

	_, err := hasher.Write(b.Bytes())
	if err != nil {
		panic(err)
	}
	return hasher.Sum64()
}

// Factorial computes a!
func Factorial(a int) int {
	if a <= 0 {
		panic("Negative number does no have factorial.")
	} else if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	}
	return a * Factorial(a-1)
}

// Permutation returns all permutations of the given set
func Permutation(set []int) [][]int {
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

	return perm(set, 0, len(set)-1)
}

// Combination returns all permutations of the given set
func Combination(set []int, n int) [][]int {
	var comb func(s []int, n int) (result [][]int)
	comb = func(s []int, n int) (result [][]int) {
		if n == 0 {
			return
		} else if n == 1 {
			result = make([][]int, len(s))
			for i, e := range s {
				result[i] = []int{e}
			}
			return
		}

		for i := 0; i <= len(s)-n; i++ {
			fixed, rest := s[i], s[i+1:]

			for _, cr := range comb(rest, n-1) {
				crc := make([]int, len(cr)+1)
				copy(crc[1:], cr)
				crc[0] = fixed
				result = append(result, crc)
			}
		}

		return result
	}

	return comb(set, n)
}

func SolveHanoiBFS(h Hanoi) (Move, error) {
	set := make(map[uint64]Move, 100)
	stack := make([]Move, 0, 100)

	curMove := Move{State: h, From: -1, To: -1}

	for _, m := range curMove.AllMoves() {
		hash := m.Hash()
		if _, ok := set[hash]; !ok {
			stack = append(stack, m)
		}
	}

	for len(stack) > 0 {
		curMove, stack = stack[0], stack[1:]

		if !curMove.CanMove() {
			continue
		}

		curMove.MakeMove()

		if curMove.HasWinned() {
			return curMove, nil
		}

		for _, m := range curMove.AllMoves() {
			hash := m.Hash()
			if _, ok := set[hash]; !ok {
				stack = append(stack, m)
			}
		}
	}

	if curMove.State.HasWinned() {
		return curMove, nil
	}
	return curMove, fmt.Errorf("failed to solve")
}

// RecoverPath recovers the path as list in the correct order
func RecoverPath(lastMove *Move) []Move {
	result := make([]Move, 0, 10)

	for lastMove != nil {
		result = append(result, *lastMove)
		lastMove = lastMove.ParentState
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
