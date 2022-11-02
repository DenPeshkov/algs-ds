// Package sort contains implementations of various sorting algorithms.
package sort

import (
	"github.com/DenPeshkov/algs-ds/utils"
)

// Insertion implements insertion sort.
func Insertion[T any](arr []T, cmp utils.Comparator[T]) {
	N := len(arr)

	for i := 1; i < N; i++ {
		x := arr[i]
		j := i - 1

		for j >= 0 && cmp(arr[j], x) > 0 {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = x
	}
}
