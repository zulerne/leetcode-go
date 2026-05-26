// https://leetcode.com/problems/network-delay-time/description/
package networkdelaytime

import (
	"container/heap"
)

type NodeTime struct {
	node, dist int
}

type MinNodeHeap []NodeTime

// Len implements [heap.Interface].
func (m MinNodeHeap) Len() int {
	return len(m)
}

// Less implements [heap.Interface].
func (m MinNodeHeap) Less(i int, j int) bool {
	return m[i].dist < m[j].dist
}

// Pop implements [heap.Interface].
func (m *MinNodeHeap) Pop() any {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}

// Push implements [heap.Interface].
func (m *MinNodeHeap) Push(x any) {
	*m = append(*m, x.(NodeTime))
}

// Swap implements [heap.Interface].
func (m MinNodeHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[int][]NodeTime)
	for _, v := range times {
		node := v[0]
		neigh := NodeTime{node: v[1], dist: v[2]}
		graph[node] = append(graph[node], neigh)
	}

	dist := make(map[int]int, n)
	dist[k] = 0

	h := &MinNodeHeap{NodeTime{node: k, dist: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		nt := heap.Pop(h).(NodeTime)
		node, nodeTime := nt.node, nt.dist

		if oldNodeTime, ok := dist[node]; ok && oldNodeTime < nodeTime {
			continue
		}

		for _, v := range graph[node] {
			neigh, newNeighTime := v.node, v.dist+nodeTime
			oldNeighTime, ok := dist[neigh]

			if !ok || newNeighTime < oldNeighTime {
				dist[neigh] = newNeighTime
				heap.Push(h, NodeTime{node: neigh, dist: newNeighTime})
			}
		}
	}

	if len(dist) < n {
		return -1
	}

	maxDist := 0
	for _, v := range dist {
		maxDist = max(maxDist, v)
	}
	return maxDist
}
