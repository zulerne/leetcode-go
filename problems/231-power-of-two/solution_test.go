package poweroftwo

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
			arg:  1,
			want: true,
		},
		{
			name: "Example 2",
			arg:  16,
			want: true,
		},
		{
			name: "Example 3",
			arg:  3,
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isPowerOfTwo(tc.arg)

			kit.AssertEqual(t, got, tc.want)
		})
	}
}
