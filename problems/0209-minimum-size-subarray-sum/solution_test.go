package minimumsizesubarraysum

import (
	"fmt"
	"testing"

	_ "github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Example 1",
			arg:  "arg",
			want: "want",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// got := minSubArrayLen(4, []int{1, 4, 4})
			// got := minSubArrayLen(11, []int{1, 2, 3, 4, 5})
			got := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})

			fmt.Printf("got: %v\n", got)

			// kit.AssertEqual(t, got, tc.want)
		})
	}
}

// func Benchmark(b *testing.B) {
// 	for b.Loop() {
// 		function("arg")
// 	}
// }
