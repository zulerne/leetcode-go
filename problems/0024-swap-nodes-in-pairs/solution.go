// https://leetcode.com/problems/swap-nodes-in-pairs/description/
package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		first, second := cur.Next, cur.Next.Next
		cur.Next = second
		first.Next = second.Next
		second.Next = first
		cur = first
	}
	return dummy.Next
}
