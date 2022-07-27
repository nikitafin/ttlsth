package pkg

import "testing"

func TestListNodeFromSlice(t *testing.T) {
	input := []int{1}
	res := ListNodeFromSlice(input)
	if res.Val != 1 || res.Next != nil {
		t.Errorf("Invalid ListNode generated: %v", res)
	}

	input1 := []int{}
	res1 := ListNodeFromSlice(input1)
	if res1 != nil {
		t.Errorf("Invalid ListNode generated: %v", res1)
	}

	input2 := []int{2, 4, 3}
	res2 := ListNodeFromSlice(input2)
	if res2.Val != 2 || res2.Next.Val != 4 || res2.Next.Next.Val != 3 {
		t.Errorf("Invalid ListNode generated: %v", res2)
	}
}

func TestListNodesEqual(t *testing.T) {
	lhs := ListNodeFromSlice([]int{1, 2, 3})
	rhs := ListNodeFromSlice([]int{1, 2, 3})
	if !ListNodesEqual(lhs, rhs) {
		t.Error("Invalid")
	}
}

func TestListNodesEqual1(t *testing.T) {
	lhs := ListNodeFromSlice([]int{1})
	rhs := ListNodeFromSlice([]int{1, 2, 3})
	if ListNodesEqual(lhs, rhs) {
		t.Error("Invalid")
	}
}

func TestListNodesEqual2(t *testing.T) {
	lhs := ListNodeFromSlice([]int{1, 2, 3})
	rhs := ListNodeFromSlice([]int{1})
	if ListNodesEqual(lhs, rhs) {
		t.Error("Invalid")
	}
}

func TestListNodesEqual3(t *testing.T) {
	lhs := ListNodeFromSlice([]int{})
	rhs := ListNodeFromSlice([]int{})
	if !ListNodesEqual(lhs, rhs) {
		t.Error("Invalid")
	}
}
