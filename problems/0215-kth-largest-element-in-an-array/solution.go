// https://leetcode.com/problems/kth-largest-element-in-an-array/description/
package kthlargestelementinanarray

import (
	"container/heap"
)

type MinHeap []int

// Len implements [heap.Interface].
func (m *MinHeap) Len() int {
	return len(*m)
}

// Less implements [heap.Interface].
func (m *MinHeap) Less(i int, j int) bool {
	return (*m)[i] < (*m)[j]
}

// Pop implements [heap.Interface].
func (m *MinHeap) Pop() any {
	old := *m
	last := old[len(old)-1]
	*m = old[:len(old)-1]
	return last
}

// Push implements [heap.Interface].
func (m *MinHeap) Push(x any) {
	old := *m
	*m = append(old, x.(int))
}

// Swap implements [heap.Interface].
func (m *MinHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func findKthLargest(nums []int, k int) int {
	h := MinHeap(nums[:k])
	heap.Init(&h)
	for _, num := range nums[k:] {
		if num > h[0] {
			heap.Pop(&h)
			heap.Push(&h, num)
		}
	}

	return h[0]
}
