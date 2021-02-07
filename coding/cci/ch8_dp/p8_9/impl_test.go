package p8_9_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_9"
)

var _ = Describe("Impl", func() {
	Context("paren comb", func() {
		It("should return empty", func() {
			result := impl.ParenthesesCombinations(0)
			Expect(result).To(HaveLen(1))
			Expect(result).To(ContainElements(""))
		})
		It("should return 1 pair", func() {
			result := impl.ParenthesesCombinations(1)
			Expect(result).To(HaveLen(1))
			Expect(result).To(ContainElements("()"))
		})
		It("should return 2 pair", func() {
			result := impl.ParenthesesCombinations(2)
			Expect(result).To(HaveLen(2))
			Expect(result).To(ContainElements(
				"()()", "(())"))
		})
		It("should return 3 pair", func() {
			result := impl.ParenthesesCombinations(3)
			Expect(result).To(HaveLen(5))
			Expect(result).To(ContainElements(
				"()()()", "((()))", "(())()", "()(())", "(()())"))
		})
	})
})
