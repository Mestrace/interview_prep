package p8_11_test

import (
	"fmt"

	impl "github.com/mestrace/interview_prep/coding/cci/ch8_dp/p8_11"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type coinsFuncType func(int) int

var _ = Describe("Impl", func() {
	testCoins := func(ctxt string, coinsFunc coinsFuncType) {
		Context(ctxt, func() {
			It("should calculate 0 cents", func() {
				result := coinsFunc(0)
				Expect(result).To(Equal(0))
			})
			It("should calculate 1 cents", func() {
				result := coinsFunc(1)
				Expect(result).To(Equal(1))
			})
			It("should calculate 2 cents", func() {
				result := coinsFunc(2)
				Expect(result).To(Equal(1))
			})
			It("should calculate 5 cents", func() {
				result := coinsFunc(5)
				Expect(result).To(Equal(2)) // 5 * 1 or 1 * 5
			})
			It("should calculate 10 cents", func() {
				result := coinsFunc(10)
				Expect(result).To(Equal(4))
				// 10
				// 5 + 5
				// 5 + 1 * 5
				// 1 * 10
			})
			It("should calculate 14 cents", func() {
				result := coinsFunc(14)
				Expect(result).To(Equal(4))
				// 10 + 1 * 4
				// 5 + 5 + 1 * 4
				// 5 + 1 * 9
				// 1 * 14
			})

			It("should calculate 15 cents", func() {
				result := coinsFunc(15)
				Expect(result).To(Equal(6))
				// 10 + 5
				// 10 + 1 * 5
				// 5 + 5 + 5
				// 5 + 5 + 1 * 5
				// 5 + 1 * 10
				// 1 * 15
			})

			It("should calculate 27 cents", func() {
				result := coinsFunc(27)
				Expect(result).To(Equal(13))
			})
		})
	}

	coinFuncToTest := map[string]func(int) int{
		"Coins":         impl.Coins,
		"CoinsMod":      impl.CoinsMod,
		"CoinsModCache": impl.CoinsModCache,
	}

	for ctxt, coinFunc := range coinFuncToTest {
		testCoins(ctxt, coinFunc)
	}

	Context("benchMark", func() {
		benchCents := func(cents int, times int) {
			Measure(fmt.Sprintf("runtime of %d", cents), func(b Benchmarker) {
				for ctxt, coinFunc := range coinFuncToTest {
					b.Time(ctxt, func() {
						coinFunc(cents)
					})
				}
			}, times)
		}

		benchCents(100, 10)
		benchCents(994, 10)
		benchCents(2028, 5)
		benchCents(3079, 3)
	})

})
