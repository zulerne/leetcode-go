package searchsuggestionssystem

import (
	"fmt"
	"testing"

	_ "github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	searchWord := "mouse"

	for i := range len(searchWord) {
		fmt.Println(searchWord[:i+1])
	}
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
			// got := function(tc.arg)

			// kit.AssertEqual(t, got, tc.want)
		})
	}
}

// func Benchmark(b *testing.B) {
// 	for b.Loop() {
// 		function("arg")
// 	}
// }
