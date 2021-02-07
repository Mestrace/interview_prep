package p8_7_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_7"
)

var _ = Describe("Impl", func() {
	Context("perm", func() {
		It("should perm 0", func() {
			Expect(impl.Permutations("")).To(HaveLen(0))
		})
		It("should perm 1", func() {
			result := impl.Permutations("C")
			Expect(result).To(HaveLen(1))
			Expect(result).To(ContainElements("C"))
		})
		It("should perm 2", func() {
			result := impl.Permutations("AC")
			Expect(result).To(HaveLen(2))
			Expect(result).To(ContainElements("AC", "CA"))
		})
		It("should perm 3", func() {
			result := impl.Permutations("ABC")
			Expect(result).To(HaveLen(6))
			Expect(result).To(ContainElements("ABC", "ACB", "BAC", "BCA", "CAB", "CBA"))
		})
	})
})
