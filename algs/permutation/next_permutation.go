package permutation

import (
	"github.com/DenPeshkov/algs-ds/utils"
)

/*
Next permutes the input slice into the next permutation in lexicographic order.
If such permutation doesn't exist the input is unchanged.
Returns true if next permutation exists, false otherwise.
*/
func Next[T any](x []T, cmp utils.Comparator[T]) bool {
	N := len(x)

	var k, l int
	notLast := false

	// largest k such that x[k]<x[k+1]a
	for k = N - 2; k >= 0; k-- {
		if cmp(x[k], x[k+1]) < 0 {
			break
		}
	}

	// not last permutation
	if k >= 0 {
		notLast = true

		// largest l>k such that x[k]<x[l]
		for l = N - 1; l > k; l-- {
			if cmp(x[k], x[l]) < 0 {
				break
			}
		}

		x[k], x[l] = x[l], x[k]
	}

	reverse(x[k+1:])

	return notLast
}

func reverse[T any](x []T) {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}
