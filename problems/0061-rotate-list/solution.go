// https://leetcode.com/problems/rotate-list/description/
package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 12345
// 1 51234
// 2 45123
// 3 34512
// 4 23451
// 5 12345
// 6 51234
//
// 12
// 1 21
// 2 12
// 3 21
// 4 12
// 5 21

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	last := head
	for last.Next != nil {
		last = last.Next
		n++
	}

	k = k % n
	if k == 0 {
		return head
	}

	last.Next = head

	newEnd := last

	for i := 0; i < n-k; i++ {
		newEnd = newEnd.Next
	}
	newHead := newEnd.Next
	newEnd.Next = nil

	return newHead
}
