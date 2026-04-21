package reversevowels

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Example 1",
			str:  "IceCreAm",
			want: "AceCreIm",
		},
		{
			name: "Example 2",
			str:  "leetcode",
			want: "leotcede",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := reverseVowels(tc.str)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
