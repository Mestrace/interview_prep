/*
Package p8_11 solves the following problem:

Coins: Given an infinite number of quarters (25 cents), dimes (10 cents), nickles (5 cent),
and pennies (1 cent), write code to calculate the number of ways of representing
n cents.
*/

package p8_11

var INTERGER_MAX = int(^uint(0) >> 1)

var ALLOWED_COINS_SORTED = []int{25, 10, 5, 1}

func Coins(n int) int {
	if n < 0 {
		panic("cannot have negative amount")
	} else if n == 0 {
		return 0
	}
	return coins(n, INTERGER_MAX)
}

func coins(n int, maxCoinAmt int) int {
	if n == 0 {
		return 1
	}
	nextAmtIdx := 0
	for ; nextAmtIdx < len(ALLOWED_COINS_SORTED); nextAmtIdx++ {
		if ALLOWED_COINS_SORTED[nextAmtIdx] <= maxCoinAmt {
			break
		}
	}

	numWays := 0
	for i := nextAmtIdx; i < len(ALLOWED_COINS_SORTED); i++ {
		if ALLOWED_COINS_SORTED[i] <= n {
			numWays += coins(n-ALLOWED_COINS_SORTED[i], ALLOWED_COINS_SORTED[i])
		}
	}

	return numWays
}

func CoinsMod(n int) int {
	if n < 0 {
		panic("cannot have negative amount")
	} else if n == 0 {
		return 0
	}
	return coinsMod(n, 0)
}

func coinsMod(n int, nextAmtIdx int) int {
	if n == 0 {
		return 1
	} else if nextAmtIdx >= len(ALLOWED_COINS_SORTED) {
		return 0
	}

	numWays := 0
	coinAmt := ALLOWED_COINS_SORTED[nextAmtIdx]

	if coinAmt > n {
		return coinsMod(n, nextAmtIdx+1)
	}
	times := n / coinAmt
	for t := times; t >= 0; t-- {
		numWays += coinsMod(n-coinAmt*t, nextAmtIdx+1)
	}

	return numWays
}

func CoinsModCache(n int) int {
	if n < 0 {
		panic("cannot have negative amount")
	} else if n == 0 {
		return 0
	}

	type coinsModParams struct {
		N          int
		NextAmtIdx int
	}

	cache := make(map[coinsModParams]int, 10)

	coinsModCache := func(n int, nextAmtIdx int) int {
		if n == 0 {
			return 1
		} else if nextAmtIdx >= len(ALLOWED_COINS_SORTED) {
			return 0
		}

		p := coinsModParams{n, nextAmtIdx}
		if result, ok := cache[p]; ok {
			return result
		}

		numWays := 0
		coinAmt := ALLOWED_COINS_SORTED[nextAmtIdx]

		if coinAmt > n {
			return coinsMod(n, nextAmtIdx+1)
		}
		times := n / coinAmt
		for t := times; t >= 0; t-- {
			numWays += coinsMod(n-coinAmt*t, nextAmtIdx+1)
		}

		cache[p] = numWays
		return cache[p]
	}

	return coinsModCache(n, 0)
}
