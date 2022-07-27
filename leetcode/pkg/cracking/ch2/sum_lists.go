package ch2

import (
	"leetcode/pkg"
)

// To slow
func AddTwoNumbers1(l1 *pkg.ListNode, l2 *pkg.ListNode) *pkg.ListNode {
	var remainder int
	var head, current *pkg.ListNode

	if l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + remainder
		remainder = sum / 10
		head = &pkg.ListNode{sum % 10, nil}
		l1 = l1.Next
		l2 = l2.Next
	} else if l1 != nil {
		head = &pkg.ListNode{l1.Val, nil}
		l1 = l1.Next
	} else {
		head = &pkg.ListNode{l2.Val, nil}
		l2 = l2.Next
	}
	current = head

	for l1 != nil && l2 != nil {

		sum := l1.Val + l2.Val + remainder
		remainder = sum / 10

		current.Next = &pkg.ListNode{sum % 10, nil}
		current = current.Next

		l1, l2 = l1.Next, l2.Next
	}

	for l1 != nil {
		sum := l1.Val + remainder
		remainder = sum / 10

		current.Next = &pkg.ListNode{sum % 10, nil}
		current = current.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + remainder
		remainder = sum / 10

		current.Next = &pkg.ListNode{sum % 10, nil}
		current = current.Next
		l2 = l2.Next
	}

	if remainder != 0 {
		current.Next = &pkg.ListNode{remainder, nil}
	}

	return head
}

func AddTwoNumbers(l1 *pkg.ListNode, l2 *pkg.ListNode) *pkg.ListNode {
	var head, current *pkg.ListNode
	var sum, remainder int

	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + remainder
		remainder = sum / 10

		if head != nil {
			current.Next = &pkg.ListNode{sum % 10, nil}
			current = current.Next
		} else {
			head = &pkg.ListNode{sum % 10, nil}
			current = head
		}

		l1, l2 = l1.Next, l2.Next
	}

	for l1 != nil {
		sum := l1.Val + remainder
		remainder = sum / 10

		current.Next = &pkg.ListNode{sum % 10, nil}
		current = current.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + remainder
		remainder = sum / 10

		current.Next = &pkg.ListNode{sum % 10, nil}
		current = current.Next
		l2 = l2.Next
	}

	if remainder != 0 {
		current.Next = &pkg.ListNode{remainder, nil}
	}

	return head
}
