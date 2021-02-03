/*
Package p8_4 solves the following problem

Write a method to return all subsets of a set
*/
package p8_4

import (
	"reflect"
	"sort"
)

// Set is a list of integers for simplicity
type Set []int

// PowerSet finds all subsets of the given set
func PowerSet(set Set) (subsets []Set) {
	if len(set) == 0 {
		return []Set{{}} // An empty set
	}

	for _, subsubset := range PowerSet(set[0 : len(set)-1]) {
		subsets = append(subsets, subsubset)

		subsubsetCopy := copySlice(subsubset).(Set)
		subsubsetCopy = append(subsubsetCopy, set[len(set)-1])

		sort.Slice(subsubsetCopy, func(i, j int) bool {
			return subsubsetCopy[i] < subsubsetCopy[j]
		})

		subsets = append(subsets, subsubsetCopy)
	}

	sort.Slice(subsets, func(i, j int) bool {
		seti := subsets[i]
		setj := subsets[j]

		if len(seti) != len(setj) {
			return len(seti) < len(setj)
		}

		for k := range seti {
			if seti[k] < setj[k] {
				return true
			}
		}
		return false
	})

	return subsets
}

// copySlice returns a shallow copy of a slice
func copySlice(s interface{}) interface{} {
	t, v := reflect.TypeOf(s), reflect.ValueOf(s)
	c := reflect.MakeSlice(t, v.Len(), v.Len())
	reflect.Copy(c, v)
	return c.Interface()
}
