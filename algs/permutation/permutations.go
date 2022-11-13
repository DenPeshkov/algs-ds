// Package permutation provides implementations of algorithms generating permutations of the input.
package permutation

// Backtracking generates all permutations of the input slice using backtracking.
func Backtracking[T any](x []T) [][]T {
	var perms [][]T
	permTemp := make([]T, len(x))

	backtrack(x, permTemp, -1, &perms)

	return perms
}

func backtrack[T any](x []T, perm []T, k int, res *[][]T) {
	// if is a solution
	if k == len(x)-1 {
		// process solution
		*res = append(*res, append([]T(nil), perm...))

		return
	}

	k++

	// generate candidates and recurse
	for i := k; i < len(x); i++ {
		x[k], x[i] = x[i], x[k]

		perm[k] = x[k]

		backtrack(x, perm, k, res)

		// to return to init state
		x[k], x[i] = x[i], x[k]
	}
}
