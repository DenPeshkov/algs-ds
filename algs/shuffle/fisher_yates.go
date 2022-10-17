// Package shuffle contains implementations of shuffling algorithms.
package shuffle

import "math/rand"

// FisherYates implements "inside-out" version of Fisher-Yates shuffle.
func FisherYates[T any](arr []T) {
	for i := 1; i < len(arr); i++ {
		r := rand.Intn(i + 1)

		arr[i], arr[r] = arr[r], arr[i]
	}
}
