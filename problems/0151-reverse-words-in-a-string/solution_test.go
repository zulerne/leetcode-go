package reversewords

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Example 1",
			arg:  "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "Example 2",
			arg:  "  hello world  ",
			want: "world hello",
		},
		{
			name: "Example 3",
			arg:  "a good   example",
			want: "example good a",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := reverseWords(tc.arg)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
