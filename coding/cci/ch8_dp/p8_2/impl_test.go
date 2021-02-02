package p8_2_test

import (
	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Impl", func() {
	Context("Robot in a grid - DFS", func() {
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

	Context("Unique paths", func() {
		It("should work without obstacles", func() {
			grid := [][]uint8{
				{0, 0, 0},
				{0, 0, 0},
			}

			Expect(impl.UniquePathDP(grid)).To(Equal(3))
		})

		It("should work with 1 obstacles", func() {
			grid := [][]uint8{
				{0, 1, 0},
				{0, 0, 0},
			}

			Expect(impl.UniquePathDP(grid)).To(Equal(1))
		})

		It("should work with no path at all", func() {
			grid := [][]uint8{
				{0, 1, 0},
				{0, 1, 0},
			}

			Expect(impl.UniquePathDP(grid)).To(Equal(0))
		})

		It("should work on one column without obstacles", func() {
			grid := [][]uint8{
				{0},
				{0},
				{0},
				{0},
				{0},
				{0},
			}

			Expect(impl.UniquePathDP(grid)).To(Equal(1))
		})

		It("should work on one column without 1 obstacles", func() {
			grid := [][]uint8{
				{0},
				{0},
				{0},
				{1},
				{0},
				{0},
			}

			Expect(impl.UniquePathDP(grid)).To(Equal(0))
		})

		It("should work on one row with no obstacles", func() {
			grid := [][]uint8{
				{0, 0, 0, 0, 0, 0},
			}

			Expect(impl.UniquePathDP(grid)).To(Equal(1))
		})

		It("should work on one row with no obstacles", func() {
			grid := [][]uint8{
				{0, 0, 1, 0, 0, 0},
			}

			Expect(impl.UniquePathDP(grid)).To(Equal(0))
		})

	})
})
