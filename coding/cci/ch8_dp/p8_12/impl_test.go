package p8_12_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_12"
)

var _ = Describe("Impl", func() {
	Context("Eight-Queen basics", func() {
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
			Expect(b.CanPlaceQueen(5, 6)).To(BeTrue())
			MustPlaceQueen(b, 5, 6)
			Expect(b.CanPlaceQueen(6, 4)).To(BeTrue())
			MustPlaceQueen(b, 6, 4)
			Expect(b.CanPlaceQueen(0, 3)).To(BeTrue())
			MustPlaceQueen(b, 0, 3)
		})
	})
})

func MustPlaceQueen(b impl.EQBoard, x, y int) {
	b.PlaceQueen(x, y)
	Expect(b.Board[y][x]).To(BeNumerically(">=", 1))
	Expect(b.AttCol[y]).To(BeNumerically(">=", 1))
	Expect(b.AttRow[x]).To(BeNumerically(">=", 1))
	Expect(b.AttLeftDiag[x+y]).To(BeNumerically(">=", 1))
	Expect(b.AttRightDiag[7-(y-x)]).To(BeNumerically(">=", 1))
}
