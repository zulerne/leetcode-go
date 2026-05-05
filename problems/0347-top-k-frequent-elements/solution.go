// https://leetcode.com/problems/top-k-frequent-elements/description/
package topkfrequentelements

import (
	"container/heap"
)

// INFO: MinHeap

type NumFreq struct {
	num  int
	freq int
}

type MinHeap []NumFreq

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(NumFreq))
}

func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}

	for num, freq := range counter {
		heap.Push(h, NumFreq{num, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, h.Len())
	for i, el := range *h {
		result[i] = el.num
	}
	return result
}
