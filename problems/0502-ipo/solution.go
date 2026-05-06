// https://leetcode.com/problems/ipo/description/
package ipo

import (
	"container/heap"
)

type Project struct {
	capital int
	profit  int
}

// INFO: MaxProfitHeap
type MaxProfitHeap []Project

// Less implements [heap.Interface].
func (h *MaxProfitHeap) Less(i int, j int) bool {
	hh := *h
	return hh[i].profit > hh[j].profit
}

// Pop implements [heap.Interface].
func (h *MaxProfitHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// Push implements [heap.Interface].
func (h *MaxProfitHeap) Push(x any) {
	old := *h
	*h = append(old, x.(Project))
}

// Swap implements [heap.Interface].
func (h *MaxProfitHeap) Swap(i int, j int) {
	hh := *h
	hh[i], hh[j] = hh[j], hh[i]
}

func (h *MaxProfitHeap) Len() int {
	return len(*h)
}

// INFO: MinCapitalHeap
type MinCapitalHeap []Project

// Less implements [heap.Interface].
func (h *MinCapitalHeap) Less(i int, j int) bool {
	hh := *h
	return hh[i].capital < hh[j].capital
}

// Pop implements [heap.Interface].
func (h *MinCapitalHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// Push implements [heap.Interface].
func (h *MinCapitalHeap) Push(x any) {
	old := *h
	*h = append(old, x.(Project))
}

// Swap implements [heap.Interface].
func (h *MinCapitalHeap) Swap(i int, j int) {
	hh := *h
	hh[i], hh[j] = hh[j], hh[i]
}

func (h *MinCapitalHeap) Len() int {
	return len(*h)
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	cap := w
	profH := MaxProfitHeap{}
	capH := MinCapitalHeap{}

	heap.Init(&profH)
	heap.Init(&capH)
	for i := range profits {
		project := Project{profit: profits[i], capital: capital[i]}
		if cap >= project.capital {
			heap.Push(&profH, project)
		} else {
			heap.Push(&capH, project)
		}
	}

	for range k {
		for capH.Len() > 0 && cap >= capH[0].capital {
			heap.Push(&profH, heap.Pop(&capH))
		}
		if profH.Len() == 0 {
			break
		}
		cap += heap.Pop(&profH).(Project).profit
	}

	return cap
}
