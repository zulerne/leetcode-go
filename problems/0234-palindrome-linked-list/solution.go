// https://leetcode.com/problems/palindrome-linked-list/description/
package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func isPalindrome(head *ListNode) bool {
	mid := middle(head)
	second := reverse(mid)
	first := head

	for second != nil {
		if first.Val != second.Val {
			return false
		}
		second = second.Next
		first = first.Next
	}

	return true
}
