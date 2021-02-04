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
	if h[i][len(h[i])-1] < h[j][len(h[j])-1] {
		v := h[i][len(h[i])-1]
		h[i] = h[i][:len(h[i])-1]
		h[j] = append(h[j], v)
		return true
	}
	return false
}
