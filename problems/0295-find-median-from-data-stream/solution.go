// https://leetcode.com/problems/find-median-from-data-stream/description/
package findmedianfromdatastream

import (
	"container/heap"
)

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

type MinHeap []int

// Less implements [heap.Interface].
func (h *MinHeap) Less(i int, j int) bool {
	hh := *h
	return hh[i] < hh[j]
}

// Pop implements [heap.Interface].
func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// Push implements [heap.Interface].
func (h *MinHeap) Push(x any) {
	old := *h
	*h = append(old, x.(int))
}

// Swap implements [heap.Interface].
func (h *MinHeap) Swap(i int, j int) {
	hh := *h
	hh[i], hh[j] = hh[j], hh[i]
}

func (h *MinHeap) Len() int {
	return len(*h)
}

type MedianFinder struct {
	left  MaxHeap
	right MinHeap
}

func Constructor() MedianFinder {
	left := MaxHeap{}
	right := MinHeap{}
	heap.Init(&left)
	heap.Init(&right)
	return MedianFinder{
		left:  left,
		right: right,
	}
}

// l: 1, r: -> l: r:1 -> l:1, r:
// l: 2,1, r: -> l:1, r:2
// l: 3,1, r:2 -> l:1, r:2,3 -> l:2,1, r:3
// l: 4,2,1 r:3 -> l:2,1 r: 3,4
func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.left, num)
	heap.Push(&this.right, heap.Pop(&this.left))
	if this.left.Len() < this.right.Len() {
		heap.Push(&this.left, heap.Pop(&this.right))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64(this.left[0])
	}
	return (float64(this.left[0]) + float64(this.right[0])) / 2
}
