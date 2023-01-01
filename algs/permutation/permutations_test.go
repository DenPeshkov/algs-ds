package permutation_test

import (
	. "github.com/DenPeshkov/algs-ds/algs/permutation"
	"testing"
)

func TestBacktracking(t *testing.T) {
	testcases := [][]int{
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{2, 1, 3},
		{1, 1, 1, 1, 1},
	}

	for _, tc := range testcases {
		perms := Backtracking(tc)

		wantedPermsNum, gotPermsNum := factorial(len(tc)), len(perms)

		if wantedPermsNum != gotPermsNum {
			t.Errorf("Not all permutations created %v: wanted %v, got %v", tc, wantedPermsNum, gotPermsNum)
		}
	}
}

func factorial(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}
