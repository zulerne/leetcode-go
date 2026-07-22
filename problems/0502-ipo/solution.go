// https://leetcode.com/problems/ipo/description/
package ipo

import (
	"container/heap"
)

type Project struct {
	capital int
	profit  int
}

type MinCapitalHeap []Project

// Len implements [heap.Interface].
func (m *MinCapitalHeap) Len() int {
	return len(*m)
}

// Less implements [heap.Interface].
func (m *MinCapitalHeap) Less(i int, j int) bool {
	return (*m)[i].capital < (*m)[j].capital
}

// Pop implements [heap.Interface].
func (m *MinCapitalHeap) Pop() any {
	old := *m
	last := old[len(old)-1]
	*m = old[:len(old)-1]
	return last
}

// Push implements [heap.Interface].
func (m *MinCapitalHeap) Push(x any) {
	old := *m
	*m = append(old, x.(Project))
}

// Swap implements [heap.Interface].
func (m *MinCapitalHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

type MaxProfitHeap []Project

// Len implements [heap.Interface].
func (m *MaxProfitHeap) Len() int {
	return len(*m)
}

// Less implements [heap.Interface].
func (m *MaxProfitHeap) Less(i int, j int) bool {
	return (*m)[i].profit > (*m)[j].profit
}

// Pop implements [heap.Interface].
func (m *MaxProfitHeap) Pop() any {
	old := *m
	last := old[len(old)-1]
	*m = old[:len(old)-1]
	return last
}

// Push implements [heap.Interface].
func (m *MaxProfitHeap) Push(x any) {
	old := *m
	*m = append(old, x.(Project))
}

// Swap implements [heap.Interface].
func (m *MaxProfitHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	curCapital := w
	minCapH := MinCapitalHeap{}
	maxProfH := MaxProfitHeap{}

	heap.Init(&minCapH)
	heap.Init(&maxProfH)

	for i := range profits {
		heap.Push(&minCapH, Project{
			profit:  profits[i],
			capital: capital[i],
		})
	}

	for range k {
		for minCapH.Len() > 0 && curCapital >= minCapH[0].capital {
			heap.Push(&maxProfH, heap.Pop(&minCapH))
		}
		if maxProfH.Len() == 0 {
			break
		}

		curCapital += heap.Pop(&maxProfH).(Project).profit
	}

	return curCapital
}
