package removeduplicatesfromsortedarrayii

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
		wantK int
	}{
		{"example1", []int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}, 5},
		{"example2", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}, 7},
		{"single", []int{1}, []int{1}, 1},
		{"two same", []int{1, 1}, []int{1, 1}, 2},
		{"all same", []int{2, 2, 2, 2}, []int{2, 2}, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			k := removeDuplicates(tc.nums)
			kit.AssertEqual(t, k, tc.wantK)
			kit.AssertEqual(t, tc.nums[:k], tc.want)
		})
	}
}

// func Benchmark(b *testing.B) {
// 	for b.Loop() {
// 		removeDuplicates([]int{1, 1, 1, 2, 2, 3})
// 	}
// }
