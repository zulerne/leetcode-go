package matrixdiagonalsum

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		arg  [][]int
		want int
	}{
		{
			name: "Example 1",
			arg: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: 25,
		},
		{
			name: "Example 2",
			arg: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			want: 8,
		},
		{
			name: "Example 3",
			arg:  [][]int{{5}},
			want: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := diagonalSum(tc.arg)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
