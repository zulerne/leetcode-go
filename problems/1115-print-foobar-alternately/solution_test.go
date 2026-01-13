package foobarconc

import (
	"strings"
	"sync"
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func Test(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "Example 1",
			n:    1,
			want: "foobar",
		},
		{
			name: "Example 1",
			n:    2,
			want: "foobarfoobar",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fooBar := NewFooBar(tc.n)
			s := strings.Builder{}

			wg := sync.WaitGroup{}
			wg.Add(tc.n * 2)

			printFoo := func() {
				s.WriteString("foo")
				wg.Done()
			}
			printBar := func() {
				s.WriteString("bar")
				wg.Done()
			}

			go fooBar.Foo(printFoo)
			go fooBar.Bar(printBar)

			wg.Wait()

			kit.AssertEqual(t, s.String(), tc.want)
		})
	}
}
