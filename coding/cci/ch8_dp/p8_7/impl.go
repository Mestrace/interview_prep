/*
Package p8_7 computes all permutations of a string of unique characters
*/
package p8_7

func Permutations(str string) []string {
	var perm func(string, int, int) []string
	perm = func(s string, l, r int) []string {
		if l == r {
			return []string{string(s)}
		}

		srune := []rune(s)

		var result []string

		for i := l; i <= r; i++ {
			srune[i], srune[l] = srune[l], srune[i]
			result = append(result, perm(string(srune), l+1, r)...)
			srune[i], srune[l] = srune[l], srune[i]
		}

		return result
	}

	return perm(str, 0, len(str)-1)
}

func PermutationsDups(str string) []string {
	srune := []rune(str)
	urune := make([]rune, 0, len(srune))
	trune := make(map[rune]struct{}, len(srune))
	for _, r := range srune {
		if _, ok := trune[r]; !ok {
			trune[r] = struct{}{}
			urune = append(urune, r)
		}
	}

	return Permutations(string(urune))
}
