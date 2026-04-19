package squaresofasortedarray

import (
	"fmt"
	"testing"
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
			got := sortedSquares([]int{-4, -1, 2, 3, 10})

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
