package ch2

import "leetcode/pkg"

func IsPalindrome(head *pkg.ListNode) bool {
	cur, runner := head, head
	var rev *pkg.ListNode

	for runner != nil && runner.Next != nil {
		runner = runner.Next.Next

		temp := cur.Next
		cur.Next = rev
		rev = cur
		cur = temp
	}

	if runner != nil {
		cur = cur.Next
	}

	for cur != nil && rev != nil {
		if cur.Val != rev.Val {
			return false
		}

		cur, rev = cur.Next, rev.Next
	}

	return true
}
