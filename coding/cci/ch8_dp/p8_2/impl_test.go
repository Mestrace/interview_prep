package p8_2_test

import (
	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Impl", func() {
	Context("DFS", func() {
		It("should work", func() {
			grid := [][]uint8{
				{0, 0, 0},
				{0, 0, 0},
			}

			path := impl.RobotInAGridDFS(grid)

			Expect(path).To(Equal([][]int{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
			}))
		})

		It("should work with obstacle", func() {
			grid := [][]uint8{
				{0, 0, 1},
				{0, 0, 0},
			}

			path := impl.RobotInAGridDFS(grid)

			Expect(path).To(Equal([][]int{
				{0, 0},
				{1, 0},
				{1, 1},
				{2, 1},
			}))
		})

		It("should return nil with no obstacle", func() {
			grid := [][]uint8{
				{0, 0, 1},
				{0, 1, 0},
			}

			path := impl.RobotInAGridDFS(grid)

			Expect(path).To(Equal([][]int(nil)))
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

func PlaceObstacles(grid [][]uint8, i, j uint) {
	grid[j][i] = 1
}
