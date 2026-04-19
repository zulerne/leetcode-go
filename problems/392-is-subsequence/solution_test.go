package issubseq

import (
	"fmt"
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	fmt.Println(isSubsequence("aaaaaa", "bbaaaa"))
	//return
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "Example 2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
		{
			name: "Example 3",
			s:    "aaaaaa",
			t:    "bbaaaa",
			want: false,
		},
		{
			name: "Example 4",
			s:    "",
			t:    "bbaaaa",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isSubsequence(tc.s, tc.t)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
