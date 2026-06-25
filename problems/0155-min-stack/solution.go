// https://leetcode.com/problems/min-stack/description/
package minstack

type pair struct {
	val int
	min int
}

type MinStack struct {
	stack []pair
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]pair, 0),
	}
}

func (this *MinStack) Push(value int) {
	curMin := value
	if len(this.stack) > 0 {
		curMin = min(curMin, this.stack[len(this.stack)-1].min)
	}
	this.stack = append(this.stack, pair{val: value, min: curMin})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
