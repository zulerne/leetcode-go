// https://leetcode.com/problems/cheapest-flights-within-k-stops/description/
package cheapestflightswithinkstops

import (
	"container/heap"
)

type HeapElem struct {
	node, price, stops int
}

type MinHeap []HeapElem

// Len implements [heap.Interface].
func (m MinHeap) Len() int {
	return len(m)
}

// Less implements [heap.Interface].
func (m MinHeap) Less(i int, j int) bool {
	return m[i].price < m[j].price
}

// Pop implements [heap.Interface].
func (m *MinHeap) Pop() any {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}

// Push implements [heap.Interface].
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(HeapElem))
}

// Swap implements [heap.Interface].
func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make(map[int][]HeapElem, n)
	for _, f := range flights {
		from, to, price := f[0], f[1], f[2]
		graph[from] = append(graph[from], HeapElem{node: to, price: price})
	}
	seen := make(map[[2]int]bool)
	h := &MinHeap{HeapElem{node: src, price: 0}}
	heap.Init(h)
	for h.Len() > 0 {
		el := heap.Pop(h).(HeapElem)
		node, price, nodeStops := el.node, el.price, el.stops
		if seen[[2]int{node, nodeStops}] {
			continue
		}
		seen[[2]int{node, nodeStops}] = true
		if node == dst {
			return price
		}
		if nodeStops > k {
			continue
		}
		for _, v := range graph[node] {
			heap.Push(h, HeapElem{node: v.node, price: price + v.price, stops: nodeStops + 1})
		}
	}
	return -1
}
