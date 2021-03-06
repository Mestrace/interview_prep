package p8_12_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_12"
)

var _ = Describe("Impl", func() {
	Context("Eight-Queen basics", func() {
		// CanPlaceQueen and PlaceQueen
		It("should place 1 successfully", func() {
			b := impl.NewEQBoard()
			MustPlaceQueen(b, 7, 7)
		})
		It("should not place if index out of bound", func() {
			b := impl.NewEQBoard()
			Expect(b.CanPlaceQueen(7, 9)).To(BeFalse())
			Expect(b.CanPlaceQueen(11, 3)).To(BeFalse())
			Expect(b.CanPlaceQueen(88, 99)).To(BeFalse())

		})
		It("should not place if two attacks", func() {
			b := impl.NewEQBoard()
			MustPlaceQueen(b, 5, 6)
			Expect(b.CanPlaceQueen(7, 6)).To(BeFalse()) // Row attacks
			Expect(b.CanPlaceQueen(5, 3)).To(BeFalse()) // Column attacks
			Expect(b.CanPlaceQueen(6, 7)).To(BeFalse()) // Left diagonal attacks
			Expect(b.CanPlaceQueen(7, 4)).To(BeFalse()) // Right diagonal attacks
		})
		It("should be place more successfully", func() {
			b := impl.NewEQBoard()
			MustPlaceQueen(b, 7, 7)
			MustPlaceQueen(b, 5, 6)
			MustPlaceQueen(b, 6, 4)
			MustPlaceQueen(b, 0, 3)
		})

		It("should undo all place queen", func() {
			b := impl.NewEQBoard()
			MustPlaceQueen(b, 7, 7)
			MustPlaceQueen(b, 5, 6)
			MustPlaceQueen(b, 6, 4)
			MustPlaceQueen(b, 0, 3)

			b.UndoPlaceQueen(6, 4)
			b.UndoPlaceQueen(7, 7)
			b.UndoPlaceQueen(5, 6)
			b.UndoPlaceQueen(0, 3)
			Expect(b).To(Equal(impl.NewEQBoard()))
		})

		It("should undo except 2", func() {
			b := impl.NewEQBoard()
			MustPlaceQueen(b, 7, 7)
			MustPlaceQueen(b, 5, 6)
			MustPlaceQueen(b, 6, 4)
			MustPlaceQueen(b, 0, 3)

			b.UndoPlaceQueen(5, 6)
			b.UndoPlaceQueen(0, 3)

			bn := impl.NewEQBoard()
			MustPlaceQueen(bn, 6, 4)
			MustPlaceQueen(bn, 7, 7)
			Expect(b).To(Equal(bn))
		})
	})
	Context("Eight-Queen helper functions", func() {
		// Adjacent Points
		It("should compute the adjacent points of a single point", func() {
			b := impl.Box{4, 4, 4, 4}
			result := b.AdjPoints()
			Expect(result).To(HaveLen(8))
			Expect(result).To(ContainElements(
				impl.Point{3, 3},
				impl.Point{3, 4},
				impl.Point{3, 5},
				impl.Point{4, 3},
				impl.Point{4, 5},
				impl.Point{5, 3},
				impl.Point{5, 4},
				impl.Point{5, 5},
			))
		})
		It("should compute adjacent points of a box", func() {
			b := impl.Box{5, 9, 7, 10}
			result := b.AdjPoints()
			Expect(result).To(HaveLen(22))
			Expect(result).To(ContainElements(
				impl.Point{4, 6},
				impl.Point{5, 6},
				impl.Point{6, 6},
				impl.Point{7, 6},
				impl.Point{8, 6},
				impl.Point{9, 6},
				impl.Point{10, 6},
				impl.Point{10, 7},
				impl.Point{10, 8},
				impl.Point{10, 9},
				impl.Point{10, 10},
				impl.Point{10, 11},
				impl.Point{9, 11},
				impl.Point{8, 11},
				impl.Point{7, 11},
				impl.Point{6, 11},
				impl.Point{5, 11},
				impl.Point{4, 11},
				impl.Point{4, 10},
				impl.Point{4, 9},
				impl.Point{4, 8},
				impl.Point{4, 7},
			))
		})
		// Add Point
		It("should add an existing point to the box", func() {
			b := impl.Box{3, 5, 3, 5}
			Expect(b.AddPoint(impl.Point{3, 4})).To(Equal(impl.Box{3, 5, 3, 5}))
		})
		It("should add an new point", func() {
			b := impl.Box{3, 5, 3, 5}
			Expect(b.AddPoint(impl.Point{6, 6})).To(Equal(impl.Box{3, 6, 3, 6}))
		})
	})

	Context("Solve Eight-Queen Problem", func() {
		It("should work", func() {
			b, ok := impl.SolveEightQueenDFS(4, 3)
			Expect(ok).To(BeTrue())
			impl.PrintBoard(b)
		})
	})
})

func MustPlaceQueen(b impl.EQBoard, x, y int) {
	Expect(b.CanPlaceQueen(x, y)).To(BeTrue())
	b.PlaceQueen(x, y)
	Expect(b.Board[y][x]).To(BeNumerically(">=", 1))
	Expect(b.AttCol[y]).To(BeNumerically(">=", 1))
	Expect(b.AttRow[x]).To(BeNumerically(">=", 1))
	Expect(b.AttLeftDiag[x+y]).To(BeNumerically(">=", 1))
	Expect(b.AttRightDiag[7-(y-x)]).To(BeNumerically(">=", 1))
}
