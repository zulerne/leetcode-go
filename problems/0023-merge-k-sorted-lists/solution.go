// https://leetcode.com/problems/merge-k-sorted-lists/description/
package mergeksortedlists

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func mergeKListsHeap(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	dummy := &ListNode{}
	cur := dummy
	for h.Len() > 0 {
		next := heap.Pop(h).(*ListNode)
		cur.Next = next
		cur = next

		if next.Next != nil {
			heap.Push(h, next.Next)
		}
	}

	return dummy.Next
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		}
		tail = tail.Next
	}
	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2

	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return merge(left, right)
}
