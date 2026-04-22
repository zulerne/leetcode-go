package designlinkedlist

import (
	"fmt"
	"testing"

	_ "github.com/zulerne/leetcode-go/kit"
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
			obj := Constructor()
			obj.Print()
			// obj.AddAtIndex(1, 0)
			// obj.Print()
			// fmt.Printf("obj.Get(0): %v\n", obj.Get(0))

			obj.AddAtTail(1)
			obj.Print()
			obj.AddAtTail(2)
			obj.Print()
			obj.AddAtTail(3)
			obj.Print()
			fmt.Printf("obj.Get(0): %v\n", obj.Get(0))

			obj.DeleteAtIndex(0)
			obj.Print()
			// obj.AddAtHead(1)
			// obj.Print()
			// obj.AddAtHead(2)
			// obj.Print()
			// obj.AddAtHead(3)
			// obj.Print()
			//
			// obj.Print()
			// param_1 := obj.Get(0)
			// fmt.Printf("param_1: %v\n", param_1)
			// obj.AddAtHead(1)
			// obj.Print()
			// obj.AddAtTail(3)
			// obj.Print()
			// obj.AddAtIndex(1, 2)
			// obj.Print()
			// obj.DeleteAtIndex(1)
			// obj.Print()

			// kit.AssertEqual(t, got, tc.want)
		})
	}
}

// func Benchmark(b *testing.B) {
// 	for b.Loop() {
// 		function("arg")
// 	}
// }
