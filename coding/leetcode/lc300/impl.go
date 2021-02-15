/*
Package lc300 solves the following problem:

Longest Increasing Subsequence: Given an integer array nums, return the length
of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some
or no elements without changing the order of the remaining elements.
For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
*/

package lc300

// LongestIncreasingSubsequenceRecursive recursively computes the longest increasing
// subsequence recursively.
func LongestIncreasingSubsequenceRecursive(seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	nseq, _ := longestIncreasingSubsequenceRecursive(seq)
	return nseq
}

func longestIncreasingSubsequenceRecursive(seq []int) (int, int) {
	if len(seq) == 1 {
		return 1, seq[0]
	}

	cur, seq := seq[len(seq)-1], seq[:len(seq)-1]
	nseq, rmax := longestIncreasingSubsequenceRecursive(seq)
	if cur > rmax {
		return nseq + 1, cur
	}
	return nseq, rmax
}
