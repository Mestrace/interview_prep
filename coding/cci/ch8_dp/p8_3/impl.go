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

import "fmt"

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

// LC 724 Find Pivot Index
// Given an array of integers nums, calculate the pivot index of this array.
// The pivot index is the index where the sum of all the numbers strictly to
// the left of the index is equal to the sum of all the numbers strictly to the
// index's right.
// If the index is on the left edge of the array, then the left sum is 0
// because there are no elements to the left. This also applies to the right
// edge of the array.
// Return the leftmost pivot index. If no such index exists, return -1.

// PivotIndex solves LC 724
func PivotIndex(A []int) int {
	leftsum := make([]int, len(A)+1)
	leftsum[1] = A[0]
	for i := 2; i <= len(A); i++ {
		leftsum[i] = leftsum[i-1] + A[i-1]
	}

	rightsum := make([]int, len(A)+1)
	rightsum[len(A)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		rightsum[i] = rightsum[i+1] + A[i]
	}

	fmt.Printf("%+v\n", A)
	fmt.Printf("%+v\n", leftsum)
	fmt.Printf("%+v\n", rightsum)

	for i := range A {
		if leftsum[i] == rightsum[i+1] {
			return i
		}
	}

	return -1
}
