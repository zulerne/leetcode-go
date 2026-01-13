package tripletsubseq

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	//fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	//fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	//fmt.Println(increasingTriplet([]int{1, 5, 0, 4, 1, 3}))
	//return

	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "Example 2",
			nums: []int{5, 4, 3, 2, 1},
			want: false,
		},
		{
			name: "Example 3",
			nums: []int{2, 1, 5, 0, 4, 6},
			want: true,
		},
		{
			name: "Example 4",
			nums: []int{20, 100, 10, 12, 5, 13},
			want: true,
		},
		{
			name: "Example 5",
			nums: []int{1, 2, 1, 3},
			want: true,
		},
		{
			name: "Example 6",
			nums: []int{1, 5, 0, 4, 1, 3},
			want: true,
		},
		{
			name: "Example 7",
			nums: []int{2, 5, 1, 6},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := increasingTriplet(tc.nums)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
