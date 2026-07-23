// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/
package findkpairswithsmallestsums

import (
	"container/heap"
)

type Pair struct {
	i, j int
	sum  int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	h := MinHeap{}
	heap.Init(&h)

	// первая колонка
	for i := 0; i < min(k, len(nums1)); i++ {
		heap.Push(&h, Pair{
			i:   i,
			j:   0,
			sum: nums1[i] + nums2[0],
		})
	}

	res := make([][]int, 0, k)

	for h.Len() > 0 && len(res) < k {
		p := heap.Pop(&h).(Pair)

		res = append(res, []int{
			nums1[p.i],
			nums2[p.j],
		})

		if p.j+1 < len(nums2) {
			heap.Push(&h, Pair{
				i:   p.i,
				j:   p.j + 1,
				sum: nums1[p.i] + nums2[p.j+1],
			})
		}
	}

	return res
}
