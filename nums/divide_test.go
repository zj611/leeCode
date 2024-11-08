package nums

import (
	"fmt"
	"math"
	"testing"
)

// 给定两个整数，被除数 dividend 和除数 divisor。
// 将两数相除，要求不使用乘法、除法和 mod 运算符。返回被除 数 dividend 除以除数 divisor 得到的商。

// 解法二 非递归版的二分搜索
func divide(divided int, divisor int) int {
	if divided == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	result := 0
	// 处理符号
	sign := -1
	if divided > 0 && divisor > 0 || divided < 0 && divisor < 0 {
		sign = 1
	}
	dvd, dvs := abs(divided), abs(divisor)

	// 这一题可以用二分搜索来做。要求除法运算之后的商，把商作为要搜索的目标。
	//商的取值范围是 [0, divided]，所以从 0 到被除数之间搜索。
	//利用二分，找到(商 + 1 ) * 除数 > 被除数并且 商 * 除数 ≤ 被除数 或 者 (商+1)* 除数 ≥ 被除数并且商 * 除数 < 被除数的时候，就算找到了商，
	//其余情况继续二分即可。最后还要 注意符号和题目规定的 Int32 取值范围。
	for dvd >= dvs {
		temp := dvs
		m := 1
		for temp<<1 <= dvd {
			temp <<= 1 // 左移动1位，相当于乘2，反之，右移1位相当于除以2
			m <<= 1
		}
		dvd -= temp
		result += m
		fmt.Println("temp", temp, "dvd", dvd, "m", m, "result", result)
	}
	return sign * result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestDivide(t *testing.T) {
	res := divide(29, 3)
	fmt.Println(res)

	//t1 := 3
	//
	//for t1 < 1000 {
	//	t1 <<= 1
	//	fmt.Println(t1)
	//}
}
