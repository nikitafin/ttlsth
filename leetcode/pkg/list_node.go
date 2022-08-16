package pkg

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	return fmt.Sprint(l.Val)
}

func ListNodeFromSlice(input []int) *ListNode {
	var head *ListNode
	switch len(input) {
	case 0:
		return nil
	case 1:
		head = &ListNode{input[0], nil}
		return head
	default:
		head = &ListNode{input[0], nil}
	}

	current := head
	for _, v := range input[1:] {
		current.Next = &ListNode{v, nil}
		current = current.Next
	}

	return head
}

func ListNodesEqual(lhs *ListNode, rhs *ListNode) bool {
	for lhs != nil && rhs != nil {
		if lhs.Val != rhs.Val {
			return false
		}

		lhs, rhs = lhs.Next, rhs.Next
	}
	if lhs != nil || rhs != nil {
		return false
	}

	return true
}

func ReverseList(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = rev
		rev = head
		head = temp
	}
	return rev
}
