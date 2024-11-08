package nums

import (
	"fmt"
	"testing"
)

// 计算x的y次 pow(x,y)

func TestPow(t *testing.T) {
	res := pow(3, 5)
	fmt.Println(res)
}

func pow(x float64, n int) float64 {
	// 时间复杂度 O(log n),空间复杂度 O(1)
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := pow(x, n/2) // 重点在这里，不断得除以2，如果是偶数，则tmp*tmp，如果是奇数，则tmp * tmp * x
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}
