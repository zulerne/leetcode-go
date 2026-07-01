// https://leetcode.com/problems/partition-list/description/
package partitionlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lDummy, gDummy := &ListNode{}, &ListNode{}
	lessCur, greatCur := lDummy, gDummy
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = nil

		if cur.Val < x {
			lessCur.Next = cur
			lessCur = lessCur.Next
		} else {
			greatCur.Next = cur
			greatCur = greatCur.Next
		}

		cur = next
	}

	lessCur.Next = gDummy.Next

	return lDummy.Next
}
