package p8_5_test

import (
	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_5"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Impl", func() {
	testMul := func(ctxt string, mulFunc func(int, int) int) {
		Context(ctxt, func() {
			It("should multiply 0", func() {
				defer Expect(recover()).ToNot(HaveOccurred())
				Expect(mulFunc(199999, 0)).To(Equal(0))
				Expect(mulFunc(0, 199999)).To(Equal(0))
				Expect(mulFunc(-1, 0)).To(Equal(0))
				Expect(mulFunc(0, -1)).To(Equal(0))
			})

			It("should multiply 1", func() {
				defer Expect(recover()).ToNot(HaveOccurred())
				Expect(mulFunc(199999, 1)).To(Equal(199999))
				Expect(mulFunc(1, 199999)).To(Equal(199999))
				Expect(mulFunc(-1, 1)).To(Equal(-1))
				Expect(mulFunc(1, -1)).To(Equal(-1))
			})

			It("should multiply 2", func() {
				defer Expect(recover()).ToNot(HaveOccurred())
				Expect(mulFunc(18, 2)).To(Equal(36))
				Expect(mulFunc(2, 18)).To(Equal(36))
			})

			It("should multiply many", func() {
				defer Expect(recover()).ToNot(HaveOccurred())
				Expect(mulFunc(37, 1992)).To(Equal(37 * 1992))
				Expect(mulFunc(6881, 33)).To(Equal(6881 * 33))
			})

			It("should multiply negative", func() {
				defer Expect(recover()).ToNot(HaveOccurred())
				Expect(mulFunc(-1, -1)).To(Equal(1))
				Expect(mulFunc(-1, -1)).To(Equal(1))
				Expect(mulFunc(-2, -5)).To(Equal(10))
				Expect(mulFunc(-7, 9)).To(Equal(-63))
			})
		})
	}

	testAdd := func(ctxt string, addFunc func(int, int) int) {
		Context(ctxt, func() {

			It("should add zero times", func() {
				defer Expect(recover()).ToNot(HaveOccurred())
				Expect(addFunc(1, 0)).To(Equal(0))
			})

			It("should add 1 times", func() {
				defer Expect(recover()).ToNot(HaveOccurred())
				Expect(addFunc(9, 1)).To(Equal(9))
			})

			It("should add many times", func() {
				defer Expect(recover()).ToNot(HaveOccurred())
				Expect(addFunc(13, 6)).To(Equal(13 * 6))
				Expect(addFunc(13, 7)).To(Equal(13 * 7))
				Expect(addFunc(1399, 145)).To(Equal(1399 * 145))
			})
		})
	}

	testAdd("Recursive Add", impl.Add)
	testMul("Recursive Add", impl.RecursiveMultiplyAdd)

	testAdd("Bitwise", impl.Add2)
	testMul("Bitwise", impl.RecursiveMultiplyBitwise)
})
