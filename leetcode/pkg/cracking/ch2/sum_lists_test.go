package ch2

import (
	"leetcode/pkg"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := pkg.ListNodeFromSlice([]int{2, 4, 3})
	l2 := pkg.ListNodeFromSlice([]int{5, 6, 4})

	res := AddTwoNumbers(l1, l2)
	exp := pkg.ListNodeFromSlice([]int{7, 0, 8})
	if !pkg.ListNodesEqual(res, exp) {
		t.Error("invalid")
	}
}

func TestAddTwoNumbers1(t *testing.T) {
	l1 := pkg.ListNodeFromSlice([]int{0})
	l2 := pkg.ListNodeFromSlice([]int{0})

	res := AddTwoNumbers(l1, l2)
	exp := pkg.ListNodeFromSlice([]int{0})
	if !pkg.ListNodesEqual(res, exp) {
		t.Error("invalid")
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := pkg.ListNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := pkg.ListNodeFromSlice([]int{9, 9, 9, 9})

	res := AddTwoNumbers(l1, l2)
	exp := pkg.ListNodeFromSlice([]int{8, 9, 9, 9, 0, 0, 0, 1})
	if !pkg.ListNodesEqual(res, exp) {
		t.Error("invalid")
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l1 := pkg.ListNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := pkg.ListNodeFromSlice([]int{9, 9, 9, 9})

		AddTwoNumbers(l1, l2)
	}
}

func BenchmarkAddTwoNumbers1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l1 := pkg.ListNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := pkg.ListNodeFromSlice([]int{9, 9, 9, 9})

		AddTwoNumbers1(l1, l2)
	}
}
