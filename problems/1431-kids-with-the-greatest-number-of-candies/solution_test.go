package kidswithcandies

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		name         string
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			name:         "Example 1",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			want:         []bool{true, true, true, false, true},
		},
		{
			name:         "Example 2",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			want:         []bool{true, false, false, false, false},
		},
		{
			name:         "Example 3",
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			want:         []bool{true, false, true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := kidsWithCandies(tc.candies, tc.extraCandies)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
