package p8_6_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_6"
)

var _ = Describe("Impl", func() {
	Context("combination", func() {
		It("should return self", func() {
			Expect(impl.Permutation([]int{1})).To(Equal([][]int{{1}}))
		})
		It("should comb 2", func() {
			Expect(impl.Permutation([]int{1, 2})).To(And(
				ContainElement([]int{1, 2}),
				ContainElement([]int{2, 1}),
			))
		})

		It("should comb 3", func() {
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
})
