package stack

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	tmp := min(this.minStack[len(this.minStack)-1], val)
	this.minStack = append(this.minStack, tmp)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
