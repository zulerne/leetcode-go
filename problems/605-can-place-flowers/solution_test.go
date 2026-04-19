package placeflowers

import (
	"fmt"
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func TestCanPlaceFlowers(t *testing.T) {
	fmt.Println(canPlaceFlowers([]int{0, 0}, 2))

	tests := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{
			name:      "Example 1",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			name:      "Example 2",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
		{
			name:      "Example 3",
			flowerbed: []int{1, 0, 1, 0, 1, 0, 1},
			n:         1,
			want:      false,
		},
		{
			name:      "Example 4",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			name:      "Example 5",
			flowerbed: []int{0},
			n:         1,
			want:      true,
		},
		{
			name:      "Example 6",
			flowerbed: []int{0, 0},
			n:         1,
			want:      true,
		},
		{
			name:      "Example 7",
			flowerbed: []int{1},
			n:         0,
			want:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := canPlaceFlowers(tc.flowerbed, tc.n)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
