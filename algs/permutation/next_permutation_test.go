package permutation_test

import (
	. "github.com/DenPeshkov/algs-ds/algs/permutation"
	"golang.org/x/exp/constraints"
	"sort"
	"testing"
)

func cmp[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}

func TestNextExists(t *testing.T) {
	testcases := [][]int{{1, 2, 3}, {1, 3, 3}, {1, 2}, {3, 1, 2, 0, 4}}

	for _, tc := range testcases {
		if got := Next(tc, cmp[int]); got != true {
			t.Errorf("Next() = %v, want %v", got, true)
		}
	}
}

func TestNextNotExists(t *testing.T) {
	testcases := [][]int{{}, {3, 2, 1}, {1}, {4, 3, 3, 1}}

	for _, tc := range testcases {
		if got := Next(tc, cmp[int]); got != false {
			t.Errorf("Next() = %v, want %v", got, false)
		}
	}
}

func FuzzTestNext(f *testing.F) {
	testcases := []string{"", "a", "aa", "ab", "ba", "bb", "abcabccc"}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, s string) {
		r := []rune(s)

		lessDesc := func(i, j int) bool {
			return r[i] > r[j]
		}

		isSorted := sort.SliceIsSorted(r, lessDesc)
		ok := Next(r, cmp[rune])

		if ok && isSorted {
			t.Error("Next() = true, want false")
		}
		if !ok && !isSorted {
			t.Error("Next() = false, want true")
		}
	})
}
