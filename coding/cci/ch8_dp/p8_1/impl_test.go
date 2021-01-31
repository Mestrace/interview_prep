package p8_1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_1"
)

var _ = Describe("Impl", func() {
	Context("recursion", func() {
		It("should work on the base cases", func() {
			ExpectBaseCaseToBeCorrect(impl.TripleStep_Recursion)
		})
		It("should work on the additional cases", func() {
			// For 4 steps
			// 1 + 1 + 1 + 1
			// 1 + 1 + 2
			// 1 + 3
			// 2 + 1 + 1
			// 2 + 2
			// 3 + 1
			ExpectToEqualUInt32(impl.TripleStep_Recursion, 4, 6)
		})
	})

	Context("memo", func() {
		It("should work on the base cases", func() {
			ExpectBaseCaseToBeCorrect(impl.Get_TripleStep_Memo())
		})
		It("should work on the additional cases", func() {
			// For 4 steps
			// 1 + 1 + 1 + 1
			// 1 + 1 + 2
			// 1 + 3
			// 2 + 1 + 1
			// 2 + 2
			// 3 + 1
			ExpectToEqualUInt32(impl.Get_TripleStep_Memo(), 4, 6)
		})
	})
})

// ExpectBaseCaseToBeCorrect 检查BaseCase的正确性
func ExpectBaseCaseToBeCorrect(implFunc func(uint32) uint32) {
	Expect(implFunc(0)).To(Equal(uint32(0)))
	Expect(implFunc(1)).To(Equal(uint32(1)))
	Expect(implFunc(2)).To(Equal(uint32(2)))
	Expect(implFunc(3)).To(Equal(uint32(3)))
}

func ExpectToEqualUInt32(implFunc func(uint32) uint32, p, e uint32) {
	Expect(implFunc(p)).To(Equal(e))
}
