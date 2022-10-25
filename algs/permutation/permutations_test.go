/* Package permutation provides implementations of algorithms generating permutations of the input. */
package permutation_test

import (
	. "github.com/DenPeshkov/algs-ds/algs/permutation"
	"testing"
)

func TestBacktracking(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"1 element", []int{1}},
		{"2 elements", []int{1, 2}},
		{"3 elements", []int{1, 2, 3}},
		{"4 elements", []int{1, 2, 3, 4}},
		{"10 elements", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"random order", []int{2, 1, 3}},
		{"all elements same", []int{1, 1, 1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			perms := Backtracking(tt.arr)

			wantedPermsNum, gotPermsNum := fac(len(tt.arr)), len(perms)

			if wantedPermsNum != gotPermsNum {
				t.Errorf("Not all permutations created %v: wanted %v, got %v", tt.arr, wantedPermsNum, gotPermsNum)
			}
		})
	}
}

func fac(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}
