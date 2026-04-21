package movezeroes

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		numsAfter []int
	}{
		{
			name:      "Example 1",
			nums:      []int{0, 1, 0, 3, 12},
			numsAfter: []int{1, 3, 12, 0, 0},
		},
		{
			name:      "Example 2",
			nums:      []int{0},
			numsAfter: []int{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			moveZeroes(tc.nums)

			kit.AssertEqual(t, tc.nums, tc.numsAfter)
		})
	}
}
