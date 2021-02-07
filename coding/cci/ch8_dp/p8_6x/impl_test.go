package p8_6x_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_6x"
)

var _ = Describe("Impl", func() {
	Context("permutation", func() {
		It("should return self", func() {
			Expect(impl.Permutation([]int{1})).To(Equal([][]int{{1}}))
		})
		It("should perm 2", func() {
			Expect(impl.Permutation([]int{1, 2})).To(And(
				ContainElement([]int{1, 2}),
				ContainElement([]int{2, 1}),
			))
		})

		It("should perm 3", func() {
			Expect(impl.Permutation([]int{1, 2, 3})).To(And(
				ContainElement([]int{1, 2, 3}),
				ContainElement([]int{1, 3, 2}),
				ContainElement([]int{2, 3, 1}),
				ContainElement([]int{2, 1, 3}),
				ContainElement([]int{3, 1, 2}),
				ContainElement([]int{3, 2, 1}),
			))
		})
	})

	Context("combination", func() {
		It("should return self", func() {
			Expect(impl.Combination([]int{1}, 1)).To(Equal([][]int{{1}}))
		})

		It("should comb 2", func() {
			result := impl.Combination([]int{1, 2, 3, 4}, 2)
			size := impl.Factorial(4) / (impl.Factorial(2) * impl.Factorial(2))
			Expect(result).To(HaveLen(size))
			Expect(result).To(And(
				ContainElement([]int{1, 2}),
				ContainElement([]int{1, 3}),
				ContainElement([]int{1, 4}),
				ContainElement([]int{2, 3}),
				ContainElement([]int{2, 4}),
				ContainElement([]int{3, 4}),
			))
		})

		It("should comb 3", func() {
			size := impl.Factorial(7) / (impl.Factorial(3) * impl.Factorial(4))
			Expect(impl.Combination([]int{1, 2, 3, 4, 5, 6, 7}, 3)).To(HaveLen(size))
		})
	})
})
