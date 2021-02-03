package p8_3_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_3"
)

var _ = Describe("Impl", func() {
	Context("Magic Index", func() {
		It("should solve on right", func() {
			A := []int{-10, -1, 0, 1, 2, 3, 6, 8, 9}
			result := impl.MagicIndexBinarySearch(A)
			Expect(result).To(Equal(6))
		})

		It("should solve on left", func() {
			A := []int{-10, 1, 3, 5, 6, 7, 11, 12, 14}
			result := impl.MagicIndexBinarySearch(A)
			Expect(result).To(Equal(1))
		})

		It("should not exist", func() {
			A := []int{-10, -1, 0, 1, 2, 3, 7, 8, 9}
			result := impl.MagicIndexBinarySearch(A)
			Expect(result).To(Equal(-1))
		})
	})

	Context("Pivot Index", func() {
		It("should work", func() {
			A := []int{-1, 10, -11}
			result := impl.PivotIndex(A)
			Expect(result).To(Equal(-1))
		})

		It("should work again", func() {
			A := []int{-1, 10, 0, 9}
			result := impl.PivotIndex(A)
			Expect(result).To(Equal(2))
		})

		It("should find nothing", func() {
			A := []int{-1, 10, 100}
			result := impl.PivotIndex(A)
			Expect(result).To(Equal(-1))
		})

		It("should find the left one", func() {
			A := []int{-1, 0, 0, -1}
			result := impl.PivotIndex(A)
			Expect(result).To(Equal(1))
		})

	})

})
