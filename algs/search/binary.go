// Package search contains implementations of various searching algorithms
package search

import "github.com/DenPeshkov/algs-ds/utils"

/*
Binary implements binary search. It returns the index of the target in the array if it is present.
If target is not present it returns the index to insert target.
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
