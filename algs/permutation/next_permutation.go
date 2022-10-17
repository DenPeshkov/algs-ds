package permutation

import (
	"github.com/DenPeshkov/algs-ds/utils"
)

/*
Next permutes the input array into next permutation in lexicographic order.
If such permutation doesn't exist the input is unchanged.
Returns true if next permutation exists, false otherwise.
*/
func Next[T any](arr []T, cmp utils.Comparator[T]) bool {
	N := len(arr)

	var k, l int
	notLast := false

	// largest k such that arr[k]<arr[k+1]a
	for k = N - 2; k >= 0; k-- {
		if cmp(arr[k], arr[k+1]) < 0 {
			break
		}
	}

	// not last permutation
	if k >= 0 {
		notLast = true

		// largest l>k such that arr[k]<arr[l]
		for l = N - 1; l > k; l-- {
			if cmp(arr[k], arr[l]) < 0 {
				break
			}
		}

		arr[k], arr[l] = arr[l], arr[k]
	}

	reverse(arr[k+1:])

	return notLast
}

func reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
