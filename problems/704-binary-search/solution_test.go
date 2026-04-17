package binarysearch

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "Example 2",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "Example 3",
			nums:   []int{-1, 0, 3, 5, 9, 12, 20},
			target: 12,
			want:   5,
		},
		{
			name:   "Example 4",
			nums:   []int{-1, 0, 3, 5, 9, 12, 20},
			target: -1,
			want:   0,
		},
		{
			name:   "Example 5",
			nums:   []int{-1, 0, 3, 5, 9, 12, 20},
			target: 0,
			want:   1,
		},
		{
			name:   "Example 6",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := search(tc.nums, tc.target)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
