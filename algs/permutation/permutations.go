// Package permutation provides implementations of algorithms generating permutations of the input.
package permutation

// Backtracking generates all permutations of the input array using backtracking.
func Backtracking[T any](arr []T) [][]T {
	var perms [][]T
	permTemp := make([]T, len(arr))

	backtrack(arr, permTemp, -1, &perms)

	return perms
}

func backtrack[T any](arr []T, perm []T, k int, res *[][]T) {
	// if is a solution
	if k == len(arr)-1 {
		// process solution
		*res = append(*res, append([]T(nil), perm...))

		return
	}

	k++

	// generate candidates and recurse
	for i := k; i < len(arr); i++ {
		arr[k], arr[i] = arr[i], arr[k]

		perm[k] = arr[k]

		backtrack(arr, perm, k, res)

		// to return to init state
		arr[k], arr[i] = arr[i], arr[k]
	}
}
