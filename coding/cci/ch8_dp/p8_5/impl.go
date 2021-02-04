/*
Package p8_5 solves the following problem:

Recursive Multiply: Write a recursive function to multiply to positive intergers
without using the * operator. You can use addition, subtraction, and bit shifting,
but you should minimize the number of those operators.
*/
package p8_5

var (
	// RecursiveMultiplyAdd calculates multiply recursively using addition
	RecursiveMultiplyAdd = constructRecursiveMultiply(Add)

	// RecursiveMultiplyBitwise calculates multiply recursively using bitwise operators
	RecursiveMultiplyBitwise = constructRecursiveMultiply(Add2)
)

func constructRecursiveMultiply(addFunc func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		if a == 0 || b == 0 {
			return 0
		}
		var negative bool
		if a < 0 {
			negative = !negative
			a = -a
		}
		if b < 0 {
			negative = !negative
			b = -b
		}
		if a < b {
			a, b = b, a
		}

		if negative {
			return -addFunc(a, b)
		}
		return addFunc(a, b)
	}
}

func Add(a, times int) int {
	if times == 0 {
		return 0
	}
	if times == 1 {
		return a
	}
	return a + Add(a, times-1)
}

func Add2(a, times int) int {
	if times == 0 {
		return 0
	} else if times == 1 {
		return a
	} else if times == 2 {
		return a + a
	}

	// odd number
	if times&0b1 == 1 {
		return a + Add2(a, times-1)
	}
	// even number
	return Add2(a<<1, times>>1)
}
