/*
Package p8_3 solves the following problem:

Magic Index: A magic index in an array A[0...n-1] is defined to be an index
such that A[i] = i. Given a sorted array of distinct integers, write a method
to find a magic index, if one exists, in array A.

Thought 1: binary search
- A[i] > i search left partition
- A[i] < i search right partition
- A[i] = i, found
*/
package p8_3

// MagicIndexBinarySearch finds the magic index using binary search
func MagicIndexBinarySearch(A []int) int {
	if len(A) == 0 {
		return -1
	}

	left := 0
	right := len(A) - 1

	for mid := left + (right-left)/2; left < right; mid = left + (right-left)/2 {
		if A[mid] == mid {
			return mid
		} else if A[mid] > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
