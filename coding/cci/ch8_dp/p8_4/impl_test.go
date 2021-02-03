package p8_4_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_4"
)

var _ = Describe("Impl", func() {
	Context("All Subsets", func() {
		It("should give empty set", func() {
			result := impl.PowerSet(impl.Set{})
			Expect(result).To(Equal([]impl.Set{{}}))
		})

		It("should have one element", func() {
			result := impl.PowerSet(impl.Set{1})
			Expect(result).To(Equal([]impl.Set{{}, {1}}))
		})

		It("should have two element", func() {
			result := impl.PowerSet(impl.Set{1, 2})
			Expect(result).To(Equal([]impl.Set{{}, {1}, {2}, {1, 2}}))
		})

		It("should have more element", func() {
			result := impl.PowerSet(impl.Set{1, 3, 5, 6})
			Expect(result).To(Equal([]impl.Set{{},
				{1}, {3}, {5}, {6},
				{1, 3}, {1, 5}, {1, 6}, {3, 5}, {3, 6}, {5, 6},
				{1, 3, 5}, {1, 3, 6}, {1, 5, 6}, {3, 5, 6},
				{1, 3, 5, 6}}))
		})
	})
})
