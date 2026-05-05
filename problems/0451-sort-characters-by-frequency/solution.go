// https://leetcode.com/problems/sort-characters-by-frequency/description/
package sortcharactersbyfrequency

import "container/heap"

type ByteFreq struct {
	byte byte
	freq int
}

type MaxHeap []ByteFreq

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(ByteFreq))
}

func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func ss(s string) string {
	h := &MaxHeap{}
	heap.Init(h)
	// counter := make(map[byte]int)
	counter := [75]int{}
	for i := range s {
		counter[s[i]-'0']++
	}

	for num, freq := range counter {
		heap.Push(h, ByteFreq{byte(num), freq})
	}

	res := make([]byte, len(s))
	for i := 0; h.Len() > 0; {
		last := heap.Pop(h).(ByteFreq)
		for range last.freq {
			res[i] = last.byte + '0'
			i++
		}
	}

	return string(res)
}

func frequencySort(s string) string {
	h := &MaxHeap{}
	heap.Init(h)
	// counter := make(map[byte]int)
	counter := [256]int{}
	for i := range s {
		// counter[s[i]]++
		counter[s[i]]++
	}

	for num, freq := range counter {
		// heap.Push(h, ByteFreq{num, freq})
		heap.Push(h, ByteFreq{byte(num), freq})
	}

	res := make([]byte, len(s))
	for i := 0; h.Len() > 0; {
		last := heap.Pop(h).(ByteFreq)
		for range last.freq {
			res[i] = last.byte
			i++
		}
	}

	return string(res)
}
