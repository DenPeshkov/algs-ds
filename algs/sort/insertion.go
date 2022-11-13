// Package sort contains implementations of various sorting algorithms.
package sort

import (
	"github.com/DenPeshkov/algs-ds/utils"
)

// Insertion implements insertion sort.
func Insertion[T any](x []T, cmp utils.Comparator[T]) {
	N := len(x)

	for i := 1; i < N; i++ {
		k := x[i]
		j := i - 1

		for j >= 0 && cmp(x[j], k) > 0 {
			x[j+1] = x[j]
			j--
		}

		x[j+1] = k
	}
}
