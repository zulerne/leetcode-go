package validpalindrome

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "Example 1",
			arg:  "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Example 2",
			arg:  "race a car",
			want: false,
		},
		{
			name: "Example 3",
			arg:  " ",
			want: true,
		},
		{
			name: "Example 4",
			arg:  "pP",
			want: true,
		},
		{
			name: "Example 5",
			arg:  "0P",
			want: false,
		},
		{
			name: "Example 6",
			arg:  ".P",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.arg)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
