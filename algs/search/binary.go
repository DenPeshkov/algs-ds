// Package search contains implementations of various searching algorithms
package search

import (
	"github.com/DenPeshkov/algs-ds/utils"
)

/*
Binary implements binary search. It returns the index of the target in the array if it is present.
If target is not present it returns the index to insert target.
If the array contains multiple elements with the specified value, there is no guarantee which one will be found.
An array must be sorted prior to calling the function.
*/
func Binary[T any](arr []T, target T, cmp utils.Comparator[T]) int {
	lo, hi := 0, len(arr)-1

	for lo < hi {
		m := int(uint(lo+hi) >> 1)
		v := arr[m]

		if cmp(v, target) == 0 {
			return m
		} else if cmp(v, target) < 0 {
			lo = m + 1
		} else {
			hi = m - 1
		}
	}

	return lo
}

/*
BinaryPredicate uses binary search to find and return the smallest index i in array arr
at which predicate(arr[i]) is true, assuming that, predicate(arr[i]) == true implies predicate(arr[i+1]) == true.
An array must be sorted prior to calling the function.
*/
func BinaryPredicate[T any](arr []T, predicate utils.Predicate[T]) int {
	lo, hi := 0, len(arr)

	for lo < hi {
		mid := int(uint(lo+hi) >> 1)

		if predicate(arr[mid]) == true {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}
