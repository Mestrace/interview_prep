/*
Package p8_11 solves the following problem:

Coins: Given an infinite number of quarters (25 cents), dimes (10 cents), nickles (5 cent),
and pennies (1 cent), write code to calculate the number of ways of representing
n cents.
*/

package p8_11

func Coins(n int) int {
	if n == 0 {
		return 0
	}
	return coins(n)
}

func coins(n int) int {
	numWays := 0
	if n >= 25 {
		c := n
		for ; c >= 25; c -= 25 {
		}
		numWays += coins(c)
	}
	if n >= 10 {
		c := n
		for ; c >= 10; c -= 10 {
		}
		numWays += coins(c)
	}
	if n >= 5 {
		c := n
		for ; c >= 5; c -= 5 {
		}
		numWays += coins(c)
	}
	if n >= 1 {
		c := n
		for ; c >= 1; c-- {
		}
		numWays += coins(c)
	}
	if n == 0 {
		return 1
	}
	return numWays
}
