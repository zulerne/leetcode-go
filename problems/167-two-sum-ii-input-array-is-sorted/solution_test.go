package twosumiiinputarrayissorted

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Example 1",
			arg:  "arg",
			want: "want",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// got := function(tc.arg)

			// kit.AssertEqual(t, got, tc.want)
		})
	}
}
