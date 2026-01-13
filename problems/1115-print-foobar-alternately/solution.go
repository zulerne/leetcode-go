// Package foobarconc
// https://leetcode.com/problems/print-foobar-alternately/
package foobarconc

type FooBar struct {
	n       int
	fooChan chan struct{}
	barChan chan struct{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:       n,
		fooChan: make(chan struct{}, 1),
		barChan: make(chan struct{}),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	fb.fooChan <- struct{}{}

	for i := 0; i < fb.n; i++ {
		<-fb.fooChan
		printFoo()
		fb.barChan <- struct{}{}
	}
	close(fb.barChan)
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barChan
		printBar()
		fb.fooChan <- struct{}{}
	}
	close(fb.fooChan)
}
