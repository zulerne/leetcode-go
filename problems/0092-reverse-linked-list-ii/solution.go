// https://leetcode.com/problems/reverse-linked-list-ii/description/
package reverselinkedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	cur := prev.Next

	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}
