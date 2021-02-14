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

// Point represents one point indexed by X and Y, possibly in a 2-D grid
type Point struct {
	X, Y int
}

// Box represents a box bounded by four corner indices, possibly in a 2-D grid
type Box struct {
	I, J, K, L int
}

// AdjPoints returns all the adjacent points a given defined box. This function
// will not check for out-of-bound points, nor it works on undefined boxes.
func (b Box) AdjPoints() []Point {
	result := make([]Point, 0, 10)

	for i := b.I - 1; i <= b.J; i++ {
		result = append(result, Point{i, b.K - 1})
	}
	for k := b.K - 1; k <= b.L; k++ {
		result = append(result, Point{b.J + 1, k})
	}
	for j := b.J + 1; j >= b.I; j-- {
		result = append(result, Point{j, b.L + 1})
	}
	for l := b.L + 1; l >= b.K; l-- {
		result = append(result, Point{b.I - 1, l})
	}
	return result
}

// AddPoint adds an additional point to the box and returns the new box formed.
// This function will not check the for out-of-bound points or boxes.
func (b Box) AddPoint(p Point) Box {
	if p.X < b.I {
		b.I = p.X
	} else if p.X > b.J {
		b.J = p.X
	}

	if p.Y < b.K {
		b.K = p.Y
	} else if p.Y > b.L {
		b.L = p.Y
	}

	return b
}

func solveEightQueenDFS(b EQBoard, moveSeq []Point, box Box) (EQBoard, bool) {
	// base case: if the bounding box Box{I, J, K, L} is larger than the size of
	// the board, return the board and wheather the number of moves equals the
	// size of the n-queen problem
	if box.I < 0 || box.J >= len(b.Board[0]) || box.K < 0 || box.L >= len(b.Board) {
		return b, len(moveSeq) == len(b.Board)
	}

	for _, p := range box.AdjPoints() {
		if !b.CanPlaceQueen(p.X, p.Y) {
			continue
		}
		b.PlaceQueen(p.X, p.Y)
		moveSeq = append(moveSeq, p)
		if rb, ok := solveEightQueenDFS(b, moveSeq, box.AddPoint(p)); ok {
			return rb, ok
		}
		b.UndoPlaceQueen(p.X, p.Y)
		moveSeq = moveSeq[:len(moveSeq)-1]
	}

	return b, false
}
