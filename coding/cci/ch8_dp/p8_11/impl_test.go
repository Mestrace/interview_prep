package p8_11_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_11"
)

var _ = Describe("Impl", func() {
	Context("Coins", func() {
		It("should calculate 0 cents", func() {
			result := impl.Coins(0)
			Expect(result).To(Equal(0))
		})
		It("should calculate 1 cents", func() {
			result := impl.Coins(1)
			Expect(result).To(Equal(1))
		})
		It("should calculate 2 cents", func() {
			result := impl.Coins(2)
			Expect(result).To(Equal(1))
		})
		It("should calculate 5 cents", func() {
			result := impl.Coins(5)
			Expect(result).To(Equal(2)) // 5 * 1 or 1 * 5
		})
		It("should calculate 10 cents", func() {
			result := impl.Coins(10)
			Expect(result).To(Equal(4)) // 10 * 1 or 5 * 2 or 5 * 1 + 1 * 5 or 10 * 1
		})
	})
})
