package palindromnum

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want bool
	}{
		{
			name: "Example 1",
			arg:  121,
			want: true,
		},
		{
			name: "Example 2",
			arg:  -121,
			want: false,
		},
		{
			name: "Example 3",
			arg:  10,
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.arg)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}

// Benchmark example - rename to match your function
func BenchmarkFunction(b *testing.B) {
	arg := 1

	for b.Loop() {
		isPalindrome(arg)
	}
}
