package p8_2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_2"
)

var _ = Describe("Impl", func() {
	Context("DFS", func() {
		It("should work", func() {
			impl.RobotInAGrid_DFS()
		})
	})
})

func Grid(i, j uint) (grid [][]uint8) {
	grid = make([][]uint8, j)

	for idx := range grid {
		grid[idx] = make([]uint8, i)
	}

	return
}
