package mergestrings

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{
			name:  "Example 1",
			word1: "abc",
			word2: "pqr",
			want:  "apbqcr",
		},
		{
			name:  "Example 2",
			word1: "ab",
			word2: "pqrs",
			want:  "apbqrs",
		},
		{
			name:  "Example 3",
			word1: "abcd",
			word2: "pq",
			want:  "apbqcd",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := mergeAlternately(tc.word1, tc.word2)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
