package watercontainer

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Example 2",
			height: []int{1, 1},
			want:   1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxArea(tc.height)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
