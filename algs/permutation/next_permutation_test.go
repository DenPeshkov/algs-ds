package permutation_test

import (
	. "github.com/DenPeshkov/algs-ds/algs/permutation"
	"testing"
)

func TestNext(t *testing.T) {
	tests := []struct {
		name string
		x    []int
		want bool
	}{
		{"exists", []int{1, 2, 3}, true},
		{"exists", []int{1, 3, 3}, true},
		{"not exists", []int{3, 2, 1}, false},
		{"not exists", []int{1}, false},
		{"exists", []int{1, 2}, true},
		{"exists", []int{3, 1, 2, 0, 4}, true},
		{"not exists", []int{4, 3, 3, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Next(tt.x, cmpInt); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func cmpInt(a, b int) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}
