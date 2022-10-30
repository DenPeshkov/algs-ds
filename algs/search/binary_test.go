// Based on https://github.com/golang/go/blob/master/src/sort/search_test.go

package search_test

import (
	"fmt"
	. "github.com/DenPeshkov/algs-ds/algs/search"
	"github.com/DenPeshkov/algs-ds/utils"
	"testing"
)

var data = []int{0: -10, 1: -5, 2: 0, 3: 1, 4: 2, 5: 3, 6: 5, 7: 7, 8: 11, 9: 100, 10: 100, 11: 100, 12: 1000, 13: 10000}

func genArr(n int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr
}

func TestBinary(t *testing.T) {
	cmp := func(a, b int) int {
		switch {
		case a == b:
			return 0
		case a < b:
			return -1
		default:
			return 1
		}
	}

	tests := []struct {
		name   string
		arr    []int
		target int
		i      int
	}{
		{"empty 0 0", genArr(0), 0, 0},
		{"1 0 0", genArr(1), 0, 0},
		{"1 -1 0", genArr(1), 1, 0},
		{"1 1 0", genArr(1), -1, 0},
		{"2 0 0", genArr(2), 0, 0},
		{"2 1 1", genArr(2), 1, 1},
		{"2 -1 0", genArr(2), -1, 0},
		{"2 2 1", genArr(2), 2, 1},
		{"100 0", genArr(100), 0, 0},
		{"100 99", genArr(100), 99, 99},
		{"100 49", genArr(100), 49, 49},
		{"100 50", genArr(100), 50, 50},
		{"100 51", genArr(100), 51, 51},
	}

	for _, e := range tests {
		i := Binary(e.arr, e.target, cmp)

		if i != e.i {
			t.Errorf("%s: expected index %d; got %d", e.name, e.i, i)
		}
	}

	for e, v := range data {
		name := fmt.Sprintf("data%v", e)

		i := Binary(data, v, cmp)

		if data[i] != v {
			t.Errorf("%s: expected to find %d; found %d", name, v, data[i])
		}
	}
}

func TestBinaryPredicate(t *testing.T) {
	predicate := func(x int) utils.Predicate[int] {
		return func(v int) bool {
			return v >= x
		}
	}

	tests := []struct {
		name string
		arr  []int
		f    utils.Predicate[int]
		i    int
	}{
		{"empty", genArr(0), nil, 0},
		{"1 1", genArr(1), predicate(1), 1},
		{"1 true", genArr(1), func(i int) bool { return true }, 0},
		{"1 false", genArr(1), func(i int) bool { return false }, 1},
		{"1e9 991", genArr(100), predicate(91), 91},
		{"1e9 true", genArr(100), func(i int) bool { return true }, 0},
		{"1e9 false", genArr(100), func(i int) bool { return false }, 100},
		{"data -20", data, predicate(-20), 0},
		{"data -10", data, predicate(-10), 0},
		{"data -9", data, predicate(-9), 1},
		{"data -6", data, predicate(-6), 1},
		{"data -5", data, predicate(-5), 1},
		{"data 3", data, predicate(3), 5},
		{"data 11", data, predicate(11), 8},
		{"data 99", data, predicate(99), 9},
		{"data 100", data, predicate(100), 9},
		{"data 101", data, predicate(101), 12},
		{"data 10000", data, predicate(10000), 13},
		{"data 10001", data, predicate(10001), 14},
		{"descending a", genArr(7), func(i int) bool { return []int{99, 99, 59, 42, 7, 0, -1, -1}[i] <= 7 }, 4},
		{"descending 7", genArr(100), func(i int) bool { return 100-i <= 7 }, 100 - 7},
	}

	for _, e := range tests {
		i := BinaryPredicate(e.arr, e.f)

		if i != e.i {
			t.Errorf("%s: expected index %d; got %d", e.name, e.i, i)
		}
	}
}
