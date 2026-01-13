package printinorder

import (
	"strings"
	"sync"
	"testing"

	"github.com/zulerne/leetcode-go/kit"
)

func TestPrintInOrder(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{
			name: "In order execution",
			nums: []int{1, 2, 3},
			want: "firstsecondthird",
		},
		{
			name: "Out of order execution",
			nums: []int{1, 3, 2},
			want: "firstsecondthird",
		},
		{
			name: "Reverse start",
			nums: []int{3, 2, 1},
			want: "firstsecondthird",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			foo := NewFoo()

			var sb strings.Builder
			var mu sync.Mutex
			var wg sync.WaitGroup
			wg.Add(3)

			safeWrite := func(text string) {
				mu.Lock()
				defer mu.Unlock()
				sb.WriteString(text)
				wg.Done()
			}

			printFirst := func() { safeWrite("first") }
			printSecond := func() { safeWrite("second") }
			printThird := func() { safeWrite("third") }

			for _, num := range tc.nums {
				switch num {
				case 1:
					go foo.First(printFirst)
				case 2:
					go foo.Second(printSecond)
				case 3:
					go foo.Third(printThird)
				}
			}

			wg.Wait()
			kit.AssertEqual(t, sb.String(), tc.want)
		})
	}
}
