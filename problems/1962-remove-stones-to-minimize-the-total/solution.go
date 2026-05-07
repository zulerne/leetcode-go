// https://leetcode.com/problems/remove-stones-to-minimize-the-total/description/
package removestonestominimizethetotal

import (
	"container/heap"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minStoneSum(piles []int, k int) int {
	var res int
	h := MaxHeap{}

	heap.Init(&h)
	for _, s := range piles {
		res += s
		heap.Push(&h, s)
	}

	for range k {
		if h.Len() == 0 {
			break
		}

		pile := (heap.Pop(&h).(int))
		half := pile / 2

		pile -= half
		res -= half
		heap.Push(&h, pile)
	}

	return res
}
