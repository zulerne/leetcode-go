// https://leetcode.com/problems/meeting-rooms-ii/description/
package meetingroomsii

import (
	"container/heap"
	"sort"
)

type Interval struct {
	Start, End int
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// TODO: Repeat
func MinMeetingRooms(intervals []*Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	h := MinHeap{}
	heap.Init(&h)

	for _, iv := range intervals {
		if h.Len() > 0 && h[0] <= iv.Start {
			heap.Pop(&h)
		}
		heap.Push(&h, iv.End)
	}

	return h.Len()
}
