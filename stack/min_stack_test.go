package stack

import (
	"fmt"
	"math"
	"testing"
)

/*
   实现一个最小栈，使得每次增加元素或者减少元素时，总是能获取到当前栈中的最小值
*/

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// 增加元素时，与最小栈的最后一个值进行比较，
	// 当小于时，存入该val；
	// 当大于时，取原最小栈的最后一个值插入
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

func TestMinStack(t *testing.T) {
	var a = []int{11, 99, 7, 8}

	ms := Constructor()
	for _, val := range a {
		ms.Push(val)
	}

	fmt.Println(ms.minStack)
	fmt.Println(ms.stack)
	ms.Pop()
	fmt.Println(ms.minStack)
	fmt.Println(ms.stack)
	//ms.Pop()
	//fmt.Println(ms.GetMin())
}
