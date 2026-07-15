// https://leetcode.com/problems/sort-list/description/
package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var mergeSort func(head *ListNode) *ListNode
	mergeSort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		var prev *ListNode
		fast, slow := head, head
		for fast != nil && fast.Next != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next.Next
		}

		prev.Next = nil
		left := mergeSort(head)
		right := mergeSort(slow)

		dummy := &ListNode{}
		tail := dummy
		for ; left != nil && right != nil; tail = tail.Next {
			if left.Val <= right.Val {
				tail.Next = left
				left = left.Next
			} else {
				tail.Next = right
				right = right.Next
			}
		}
		if left != nil {
			tail.Next = left
		}
		if right != nil {
			tail.Next = right
		}

		return dummy.Next
	}

	return mergeSort(head)
}
