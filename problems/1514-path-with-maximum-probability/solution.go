// https://leetcode.com/problems/path-with-maximum-probability/description/
package pathwithmaximumprobability

import "container/heap"

type HeapElem struct {
	node int
	prob float64
}

type MaxHeap []HeapElem

// Len implements [heap.Interface].
func (m MaxHeap) Len() int {
	return len(m)
}

// Less implements [heap.Interface].
func (m MaxHeap) Less(i int, j int) bool {
	return m[i].prob > m[j].prob
}

// Pop implements [heap.Interface].
func (m *MaxHeap) Pop() any {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}

// Push implements [heap.Interface].
func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(HeapElem))
}

// Swap implements [heap.Interface].
func (m MaxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	probs := make(map[int]float64, n)
	probs[start_node] = 1

	graph := make(map[int][]HeapElem, n)
	for i, v := range edges {
		prob := succProb[i]
		graph[v[0]] = append(graph[v[0]], HeapElem{node: v[1], prob: prob})
		graph[v[1]] = append(graph[v[1]], HeapElem{node: v[0], prob: prob})
	}

	h := &MaxHeap{HeapElem{node: start_node, prob: 1}}
	heap.Init(h)

	for h.Len() > 0 {
		el := heap.Pop(h).(HeapElem)
		node, prob := el.node, el.prob

		if oldProb, ok := probs[node]; ok && oldProb > prob {
			continue
		}

		for _, v := range graph[node] {
			neigh, neighProb := v.node, v.prob*prob

			oldNeighProb, ok := probs[neigh]
			if !ok || neighProb > oldNeighProb {
				probs[neigh] = neighProb
				heap.Push(h, HeapElem{node: neigh, prob: neighProb})
			}
		}
	}

	return probs[end_node]
}
