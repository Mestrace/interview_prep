/*
Package printRune solves the following problem:

** Your Problem Here **
*/

package printRune

import (
	"fmt"
	"strings"
)

var (
	emptyToken rune = '#'
)

// PrintRunes prints the result ofWriteRunes
func PrintRunes(idxSet []int, token rune) {
	fmt.Println(strings.Join(WriteRunes(idxSet, token), "\n"))
}

// WriteRunes returns result and all intermediate results of replacing the empty-token
// string by the selected token.
// Example:
// 		WriteRunes([3, 1, 2, 0, 4], "x") -> []string{"#####", "###x#", "#x#x#", "#xxx#", "xxxx#", "xxxxx"}
func WriteRunes(idxSet []int, token rune) []string {
	return writeRunes(idxSet, token, strings.Repeat(string(emptyToken), len(idxSet)))
}

func writeRunes(idxSet []int, token rune, lastString string) []string {
	if len(idxSet) == 0 {
		return []string{lastString}
	}

	idx, restIdxSet := idxSet[0], idxSet[1:]

	thisRune := []rune(lastString)
	thisRune[idx] = token

	restResults := writeRunes(restIdxSet, token, string(thisRune))

	return append([]string{lastString}, restResults...)
}
