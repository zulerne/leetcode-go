package rangesumqueryimmutable

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		left  int
		right int
		want  int
	}{
		{
			name:  "Example 1",
			nums:  []int{-2, 0, 3, -5, 2, -1},
			left:  0,
			right: 2,
			want:  1,
		},
		{
			name:  "Example 2",
			nums:  []int{-2, 0, 3, -5, 2, -1},
			left:  2,
			right: 5,
			want:  -1,
		},
		{
			name:  "Example 3",
			nums:  []int{-2, 0, 3, -5, 2, -1},
			left:  0,
			right: 5,
			want:  -3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			numArray := Constructor(tc.nums)
			got := numArray.SumRange(tc.left, tc.right)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
