// Package shuffle contains implementations of shuffling algorithms.
package shuffle

import "math/rand"

// FisherYates implements "inside-out" version of Fisher-Yates shuffle.
func FisherYates[T any](x []T) {
	for i := 1; i < len(x); i++ {
		r := rand.Intn(i + 1)

		x[i], x[r] = x[r], x[i]
	}
}
