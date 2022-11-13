package sort

import "github.com/DenPeshkov/algs-ds/utils"

// Selection implements selection sort.
func Selection[T any](x []T, cmp utils.Comparator[T]) {
	N := len(x)

	for i := 0; i < N; i++ {
		minInd := minSlice(x, i, cmp)

		x[i], x[minInd] = x[minInd], x[i]
	}
}

func minSlice[T any](x []T, from int, cmp utils.Comparator[T]) int {
	minInd := from
	min := x[minInd]

	for i := from + 1; i < len(x); i++ {
		if cmp(x[i], min) < 0 {
			min = x[i]
			minInd = i
		}
	}

	return minInd
}
