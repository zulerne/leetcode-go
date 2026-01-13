// Package printinorder
// https://leetcode.com/problems/print-in-order
package printinorder

type Foo struct {
	secondCh chan struct{}
	thirdCh  chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		secondCh: make(chan struct{}, 1),
		thirdCh:  make(chan struct{}, 1),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	f.secondCh <- struct{}{}
	close(f.secondCh)
}

func (f *Foo) Second(printSecond func()) {
	<-f.secondCh
	printSecond()
	f.thirdCh <- struct{}{}
	close(f.thirdCh)
}

func (f *Foo) Third(printThird func()) {
	<-f.thirdCh
	printThird()
}
