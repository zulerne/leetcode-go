package stringcompression

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name       string
		chars      []byte
		want       int
		charsAfter []byte
	}{
		{
			name:       "Example 1",
			chars:      []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			want:       6,
			charsAfter: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			name:       "Example 2",
			chars:      []byte{'a'},
			want:       1,
			charsAfter: []byte{'a'},
		},
		{
			name:       "Example 3",
			chars:      []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			want:       4,
			charsAfter: []byte{'a', 'b', '1', '2'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := compress(tc.chars)

			kit.AssertEqual(t, got, tc.want)
			kit.AssertEqual(t, tc.chars[:got], tc.charsAfter)
		})
	}
}
