package maxofsumpairs

import (
	"fmt"
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {

	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))

	return

	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			k:    5,
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{3, 1, 3, 4, 3},
			k:    6,
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxOperations(tc.nums, tc.k)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
