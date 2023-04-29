package main

import (
	"basic_go_practicum/pkg"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := TwoSum(nums, target)
	if !pkg.Equal(res, []int{1, 0}) {
		t.Error("TwoSum invalid")
	}

	nums1 := []int{2, 7, 11, 15}
	target1 := 1

	res1 := TwoSum(nums1, target1)
	if !pkg.Equal(res1, []int{}) {
		t.Error("TwoSum invalid")
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}
