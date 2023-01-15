package ch2

import (
	"leetcode/pkg"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	sl := []int{1, 2, 2, 1}
	l := pkg.ListNodeFromSlice(sl)
	if !IsPalindrome(l) {
		t.Errorf("%v is not palindrome", sl)
	}
}

func TestIsPalindrome1(t *testing.T) {
	sl := []int{1, 2, 3, 2, 1}
	l := pkg.ListNodeFromSlice(sl)
	if !IsPalindrome(l) {
		t.Errorf("%v is not palindrome", sl)
	}
}

func TestIsPalindrome2(t *testing.T) {
	sl := []int{1, 2}
	l := pkg.ListNodeFromSlice(sl)
	if IsPalindrome(l) {
		t.Errorf("%v is not palindrome", sl)
	}
}

func TestIsPalindrome3(t *testing.T) {
	sl := []int{1, 3, 2, 4, 3, 2, 1}
	l := pkg.ListNodeFromSlice(sl)
	if IsPalindrome(l) {
		t.Errorf("%v is not palindrome", sl)
	}
}
