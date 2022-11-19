package partition_test

import (
	. "github.com/DenPeshkov/algs-ds/algs/partition"
	"testing"
)

var data = []int{0: 1, 1: 3, 2: -1, 3: 0, 4: 0, 5: 7, 6: 18, 7: 4, 8: 4, 9: 4,
	10: 18, 11: 5, 12: 14, 13: 0, 14: 23, 15: 4, 16: 9, 17: 100, 18: 10000, 19: -34, 20: -56, 21: 3}

func reverse(x []int) []int {
	rev := make([]int, len(x))
	copy(rev, x)

	N := len(rev)

	for i := range rev {
		rev[i], rev[N-i-1] = rev[N-i-1], rev[i]
	}

	return rev
}

func cmp(a int) func(int) int {
	return func(e int) int {
		switch {
		case e == a:
			return 0
		case e < a:
			return -1
		default:
			return 1
		}
	}
}

func TestThreeWay(t *testing.T) {
	testX(data, t)
	testX(reverse(data), t)
	testX([]int{}, t)
	testX(nil, t)
	testX([]int{23}, t)
	testX([]int{1, 1001}, t)
	testX([]int{1001, 1}, t)
	testData(t)
}

func testData(t *testing.T) {
	tests := []struct {
		f  func(int) int
		lt int
		gt int
	}{
		{cmp(0), 3, 5},
		{cmp(4), 9, 12},
		{cmp(-56), 0, 0},
		{cmp(10000), 21, 21},
		{cmp(-10001), 0, -1},
		{cmp(10001), 22, 21},
		{cmp(-1), 2, 2},
		{cmp(100), 20, 20},
		{cmp(12), 16, 15},
		{cmp(1000), 21, 20},
		{cmp(-50), 1, 0},
	}

	for _, test := range tests {
		lt, gt := ThreeWay(data, test.f)

		if lt != test.lt || gt != test.gt {
			t.Errorf("expected [%v, %v], got [%v, %v]", test.lt, test.gt, lt, gt)
		}
	}
}

func testX(x []int, t *testing.T) {
	tests := []struct {
		f  func(int) int
		lt int
		gt int
	}{
		{func(int) int { return -1 }, len(x), len(x) - 1},
		{func(int) int { return 0 }, 0, len(x) - 1},
		{func(int) int { return 1 }, 0, -1},
	}

	for _, test := range tests {
		lt, gt := ThreeWay(x, test.f)

		if lt != test.lt || gt != test.gt {
			t.Errorf("expected [%v, %v], got [%v, %v]", test.lt, test.gt, lt, gt)
		}
	}
}
