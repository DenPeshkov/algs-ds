// Package search contains implementations of various searching algorithms
package search

import (
	"github.com/DenPeshkov/algs-ds/utils"
)

/*
Binary implements binary search.
It returns the index of the target in the slice if it is present.
If target is not present it returns the index to insert target. The returned index is in interval: [0, len(x)]
If the slice contains multiple elements with the specified value, there is no guarantee which one will be found.
The slice must be sorted in ascending order, defined by cmp.
*/
func Binary[T any](x []T, target T, cmp utils.Comparator[T]) int {
	lo, hi := 0, len(x)-1

	for lo <= hi {
		m := int(uint(lo+hi) >> 1)
		v := x[m]

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
BinaryPredicate uses binary search to find and return the smallest index i in the slice x
at which p(x[i]) is true, assuming that, p(x[i]) == true implies p(x[i+1]) == true.
That is, BinaryPredicate requires that p is false for some (possibly empty) prefix of the slice
and then true for the (possibly empty) remainder.
If there is no such index i, BinaryPredicate returns i = n.
*/
func BinaryPredicate[T any](x []T, p utils.Predicate[T]) int {
	lo, hi := 0, len(x)

	for lo < hi {
		mid := int(uint(lo+hi) >> 1)

		if p(x[mid]) == true {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}
