/*
Package p8_9 solves the following problem:

Implement an algorithm to print all valid combinations of n pairs of parentheses
*/
package p8_9

func ParenthesesCombinations(n int) []string {

	var parenComb func([]rune, int, int) [][]rune

	parenComb = func(pc []rune, left, right int) [][]rune {
		if left == 0 && right == 0 {
			return [][]rune{pc}
		} else if right == 0 {
			return nil
		}

		var result [][]rune

		if left > 0 {
			pcc := make([]rune, len(pc)+1)
			copy(pcc, pc)
			pcc[len(pc)] = '('
			result = append(result, parenComb(pcc, left-1, right)...)
		}

		if left < right {
			pcc := make([]rune, len(pc)+1)
			copy(pcc, pc)
			pcc[len(pc)] = ')'
			result = append(result, parenComb(pcc, left, right-1)...)
		}

		return result
	}

	var result []string
	for _, sr := range parenComb([]rune{}, n, n) {
		result = append(result, string(sr))
	}

	return result
}
