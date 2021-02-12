/*
Package p8_12 solves the following problem:

Eight Queens: Write an algorithm to print all ways of arranging eight queens on
an 8x8 chess board so that none of them share the same row, column, or diagonal.
In this case, "diagonal" means all diagonals, not just the two that bisects the
board.
*/
package p8_12

import "reflect"

type EQBoard struct {
	Board        [][]int
	AttRow       []int
	AttCol       []int
	AttLeftDiag  []int
	AttRightDiag []int
}

func NewEQBoard() EQBoard {
	board := make([][]int, 8)
	for i := range board {
		board[i] = make([]int, 8)
	}

	return EQBoard{
		Board:        board,
		AttRow:       make([]int, 8),
		AttCol:       make([]int, 8),
		AttLeftDiag:  make([]int, 15),
		AttRightDiag: make([]int, 15),
	}
}

// CanPlaceQueen checks whether a queen can be placed on the selected (x, y). It
// checks Row, Column, Left Diagonal ( \ ), and Right Diagonal ( / ) for possible
// attacks.
func (b EQBoard) CanPlaceQueen(x, y int) bool {
	if x >= len(b.Board[0]) || y > len(b.Board) {
		return false
	} else if b.Board[y][x] > 0 {
		return false
	} else if b.AttCol[y] > 0 {
		return false
	} else if b.AttRow[x] > 0 {
		return false
	} else if b.AttLeftDiag[x+y] > 0 {
		return false
	} else if b.AttRightDiag[7-(y-x)] > 0 {
		return false
	}

	return true
}

// PlaceQueen place a queen on the selected position. Does not check for validity
func (b EQBoard) PlaceQueen(x, y int) {
	b.Board[y][x]++
	b.AttCol[y]++
	b.AttRow[x]++
	b.AttLeftDiag[x+y]++
	b.AttRightDiag[7-(y-x)]++
}

// UndoPlaceQueen undo-s the operation of placing a queen. Does not check for
// validity
func (b EQBoard) UndoPlaceQueen(x, y int) {
	b.Board[y][x]--
	b.AttCol[y]--
	b.AttRow[x]--
	b.AttLeftDiag[x+y]--
	b.AttRightDiag[7-(y-x)]--
}

// Copy returns a copy of the current state
func (b EQBoard) Copy() EQBoard {
	newBoard := make([][]int, len(b.Board))
	for i := range b.Board {
		newBoard[i] = make([]int, len(b.Board[i]))
		copy(newBoard[i], b.Board[i])
	}

	newAttCol := make([]int, len(b.AttCol))
	copy(newAttCol, b.AttCol)

	newAttRow := make([]int, len(b.AttRow))
	copy(newAttRow, b.AttRow)

	newAttLeftDiag := make([]int, len(b.AttLeftDiag))
	copy(newAttLeftDiag, b.AttLeftDiag)

	newAttRightDiag := make([]int, len(b.AttRightDiag))
	copy(newAttRightDiag, b.AttRightDiag)

	return EQBoard{
		Board:        newBoard,
		AttRow:       newAttRow,
		AttCol:       newAttCol,
		AttLeftDiag:  newAttLeftDiag,
		AttRightDiag: newAttRightDiag,
	}
}

// Equal compares the equality of the targeting board and whether the state are
// equal.
func (b EQBoard) Equal(ob EQBoard) bool {
	return reflect.DeepEqual(b, ob)
}
