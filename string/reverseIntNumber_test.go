package string

import (
	"fmt"
	"testing"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//注意:假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31, 2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

func reverseIntNumber(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func TestReverseIntNumber(t *testing.T) {
	fmt.Println(reverseIntNumber(789997662))
}
