package lc300_test

import (
	impl "github.com/mestrace/interview_prep/coding/leetcode/lc300"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Impl", func() {
	Context("recursive", func() {
		It("should work with 0 element", func() {
			result := impl.LongestIncreasingSubsequenceRecursive([]int{})
			Expect(result).To(Equal(0))
		})
		It("should work with 1 element", func() {
			result := impl.LongestIncreasingSubsequenceRecursive([]int{1})
			Expect(result).To(Equal(1))
		})
		It("should work with 3 element", func() {
			result := impl.LongestIncreasingSubsequenceRecursive([]int{1, 2, 3})
			Expect(result).To(Equal(3))
		})
		It("should work with 5 element", func() {
			result := impl.LongestIncreasingSubsequenceRecursive([]int{1, 4, 5, 2, 6})
			Expect(result).To(Equal(4))
		})
	})
})
