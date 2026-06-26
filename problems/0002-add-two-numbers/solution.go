// https://leetcode.com/problems/add-two-numbers/description/
package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	acc := 0

	for l1 != nil || l2 != nil || acc > 0 {
		sum := acc
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		acc = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return dummy.Next
}
