// https://leetcode.com/problems/last-stone-weight/description/
package laststoneweight

import "container/heap"

type MaxHeap []int

// Less implements [heap.Interface].
func (h *MaxHeap) Less(i int, j int) bool {
	hh := *h
	return hh[i] > hh[j]
}

// Pop implements [heap.Interface].
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// Push implements [heap.Interface].
func (h *MaxHeap) Push(x any) {
	old := *h
	*h = append(old, x.(int))
}

// Swap implements [heap.Interface].
func (h *MaxHeap) Swap(i int, j int) {
	hh := *h
	hh[i], hh[j] = hh[j], hh[i]
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(stones)
	heap.Init(&h)

	for h.Len() >= 2 {
		first := heap.Pop(&h).(int)
		second := heap.Pop(&h).(int)
		if first != second {
			heap.Push(&h, first-second)
		}
	}

	if h.Len() > 0 {
		return h[0]
	}
	return 0
}
