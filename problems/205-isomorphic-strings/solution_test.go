package isomorphicstrings

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want bool
	}{
		{
			name: "Example 1",
			arg1: "egg",
			arg2: "add",
			want: true,
		},
		{
			name: "Example 2",
			arg1: "f11",
			arg2: "b23",
			want: false,
		},
		{
			name: "Example 3",
			arg1: "paper",
			arg2: "title",
			want: true,
		},
		{
			name: "Example 4",
			arg1: "bbbaaaba",
			arg2: "aaabbbba",
			want: false,
		},
		{
			name: "Example 5",
			arg1: "badc",
			arg2: "baba",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isIsomorphic(tc.arg1, tc.arg2)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
