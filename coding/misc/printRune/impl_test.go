package printRune_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	impl "github.com/mestrace/interview_prep/coding/misc/printRune"
)

var _ = Describe("Impl", func() {
	Context("writeRune", func() {
		It("should avoid empty", func() {
			result := impl.WriteRunes([]int{}, 'x')
			Expect(result).To(HaveLen(1))
			Expect(result).To(ContainElements(""))
		})
		It("should write one element", func() {
			result := impl.WriteRunes([]int{0}, 'x')
			Expect(result).To(HaveLen(2))
			Expect(result).To(ContainElements("#", "x"))
		})
		It("should write normally", func() {
			result := impl.WriteRunes([]int{3, 1, 2, 0, 4}, 'x')
			Expect(result).To(HaveLen(6))
			Expect(result).To(Equal([]string{"#####", "###x#", "#x#x#", "#xxx#", "xxxx#", "xxxxx"}))
		})
	})
})
