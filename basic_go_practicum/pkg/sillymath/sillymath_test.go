package sillymath

import (
	"basic_go_practicum/pkg"
	"testing"
)

func TestSliceSum(t *testing.T) {
	input := []int{
		78, 27, 41, 24, 21, 63, 39, 26, 53,
		71, 39, 67, 12, 73, 50, 61, 81, 55,
		30, 12, 56, 1, 48, 26, 19, 14, 86,
		58, 13, 96, 95, 33, 31, 7, 65, 42,
		8, 46, 36, 78, 38, 34, 35, 16, 40,
		8, 38, 72, 2, 11,
	}

	res := SliceSum(input)
	exp := 2075

	if exp != res {
		t.Errorf("%d != %d", exp, res)
	}
}

func TestMapSlice(t *testing.T) {
	input := []int{1, 2, 3}
	addOne := func(in int) int {
		return in + 1
	}

	MapSlice(input, addOne)
	exp := []int{2, 3, 4}
	if !pkg.Equal(input, exp) {
		t.Errorf("%v != %v", input, exp)
	}
}

func TestFoldSlice(t *testing.T) {
	input := []int{
		78, 27, 41, 24, 21, 63, 39, 26, 53,
		71, 39, 67, 12, 73, 50, 61, 81, 55,
		30, 12, 56, 1, 48, 26, 19, 14, 86,
		58, 13, 96, 95, 33, 31, 7, 65, 42,
		8, 46, 36, 78, 38, 34, 35, 16, 40,
		8, 38, 72, 2, 11,
	}
	exp := SliceSum(input) + 1

	sum := func(lhs, rhs int) int {
		return lhs + rhs
	}

	res := FoldSlice(input, sum, 1)
	if res != exp {
		t.Errorf("%d != %d", exp, res)
	}
}
