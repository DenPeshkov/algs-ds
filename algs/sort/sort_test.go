// Based on https://github.com/golang/go/blob/master/src/sort/sort_test.go

package sort_test

import (
	"github.com/DenPeshkov/algs-ds/algs/sort"
	"math/rand"
	gosort "sort"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

var cmp = func(a, b int) int {
	switch {
	case a == b:
		return 0
	case a < b:
		return -1
	default:
		return 1
	}
}

func reverse(arr []int) {
	for lo, hi := 0, len(arr)-1; lo < hi; lo, hi = lo+1, hi-1 {
		arr[lo], arr[hi] = arr[hi], arr[lo]
	}
}

func TestInsertion(t *testing.T) {
	sortingFunc := func(arr []int) {
		sort.Insertion(arr, cmp)
	}

	testEmpty(sortingFunc)
	testData(sortingFunc, t)
	testRandom(1_000, sortingFunc, t)
	testReverseSort(sortingFunc, t)
}

func TestSelection(t *testing.T) {
	sortingFunc := func(arr []int) {
		sort.Selection(arr, cmp)
	}

	testEmpty(sortingFunc)
	testData(sortingFunc, t)
	testRandom(1_000, sortingFunc, t)
	testReverseSort(sortingFunc, t)
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
