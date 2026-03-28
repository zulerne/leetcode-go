package reversestring

import (
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want []byte
	}{
		{
			name: "Example 1",
			arg:  []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reverseString(tc.arg)

			kit.AssertEqual(t, tc.arg, tc.want)
		})
	}
}
