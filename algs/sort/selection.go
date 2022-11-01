package sort

import "github.com/DenPeshkov/algs-ds/utils"

// Selection implements selection sort.
func Selection[T any](arr []T, cmp utils.Comparator[T]) {
	N := len(arr)

	for i := 0; i < N; i++ {
		minInd := minSlice(arr, i, cmp)

		arr[i], arr[minInd] = arr[minInd], arr[i]
	}
}

func minSlice[T any](arr []T, from int, cmp utils.Comparator[T]) int {
	minInd := from
	min := arr[minInd]

	for i := from + 1; i < len(arr); i++ {
		if cmp(arr[i], min) < 0 {
			min = arr[i]
			minInd = i
		}
	}

	return minInd
}
