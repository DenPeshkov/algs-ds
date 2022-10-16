package sort

import "github.com/DenPeshkov/algs-ds/utils"

// Selection implements selection sort.
func Selection[T any](arr []T, cmp utils.Comparator[T]) {
	N := len(arr)

	for i := 0; i < N; i++ {
		minInd := minSlice(arr[i:], cmp)

		arr[i], arr[minInd] = arr[minInd], arr[i]
	}
}

func minSlice[T any](arr []T, cmp utils.Comparator[T]) (minInd int) {
	min := arr[0]

	for i, a := range arr[1:] {
		if cmp(a, min) < 0 {
			min = a
			minInd = i
		}
	}

	return minInd
}
