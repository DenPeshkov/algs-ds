// Based on https://github.com/golang/go/blob/master/src/sort/sort_test.go

package sort_test

import (
	"github.com/DenPeshkov/algs-ds/algs/sort"
	"golang.org/x/exp/constraints"
	"math/rand"
	gosort "sort"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func cmp[T constraints.Ordered](a, b T) int {
	switch {
	case a == b:
		return 0
	case a < b:
		return -1
	default:
		return 1
	}
}

func reverse(x []int) {
	for lo, hi := 0, len(x)-1; lo < hi; lo, hi = lo+1, hi-1 {
		x[lo], x[hi] = x[hi], x[lo]
	}
}

func TestInsertion(t *testing.T) {
	sortingFunc := func(x []int) {
		sort.Insertion(x, cmp[int])
	}

	testEmpty(sortingFunc)
	testData(sortingFunc, t)
	testRandom(1_000, sortingFunc, t)
	testReverseSort(sortingFunc, t)
}

func TestSelection(t *testing.T) {
	sortingFunc := func(x []int) {
		sort.Selection(x, cmp[int])
	}

	testEmpty(sortingFunc)
	testData(sortingFunc, t)
	testRandom(1_000, sortingFunc, t)
	testReverseSort(sortingFunc, t)
}

func FuzzInsertion(f *testing.F) {
	sortingFunc := func(r []rune) {
		sort.Insertion(r, cmp[rune])
	}

	fuzzSort(sortingFunc, f)
}

func FuzzSelection(f *testing.F) {
	sortingFunc := func(r []rune) {
		sort.Selection(r, cmp[rune])
	}

	fuzzSort(sortingFunc, f)
}

func testEmpty(sortingFunc func([]int)) {
	empty := make([]int, 0)
	var nilSlice []int

	sortingFunc(empty)

	sortingFunc(nilSlice)
}

func testData(sortingFunc func([]int), t *testing.T) {
	var data []int

	copy(data, ints[:])

	sortingFunc(data)

	if !gosort.IntsAreSorted(data) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func testRandom(n int, sortingFunc func([]int), t *testing.T) {
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}

	if gosort.IntsAreSorted(data) {
		t.Fatalf("terrible rand.rand")
	}

	sortingFunc(data)

	if !gosort.IntsAreSorted(data) {
		t.Errorf("sort didn't sort - 1M ints")
	}
}

func testReverseSort(sortingFunc func(arr []int), t *testing.T) {
	var data, revData []int

	copy(data, ints[:])
	copy(revData, ints[:])
	reverse(revData)

	sortingFunc(data)
	sortingFunc(revData)

	for i := 0; i < len(data); i++ {
		if data[i] != revData[len(data)-1-i] {
			t.Errorf("reverse sort didn't sort")
		}
		if i > len(data)/2 {
			break
		}
	}
}

func fuzzSort(sortingFunc func(r []rune), f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		r := []rune(s)

		sortingFunc(r)

		if !gosort.SliceIsSorted(r, func(i, j int) bool { return r[i] < r[j] }) {
			t.Errorf("slice was not sorted")
		}
	})
}
