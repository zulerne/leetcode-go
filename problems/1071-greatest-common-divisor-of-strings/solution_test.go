package stringsdivisor

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{
			name: "Example 1",
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			name: "Example 2",
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			name: "Example 3",
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
		{
			name: "Example 4",
			str1: "AAAAAB",
			str2: "AAA",
			want: "",
		},
		{
			name: "fifth",
			str1: "ABABABAB",
			str2: "ABAB",
			want: "ABAB",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := gcdOfStrings(tc.str1, tc.str2)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
