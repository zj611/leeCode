package dp

import (
	"fmt"
	"testing"
)

//最大连续子序列和

func maxSumNums(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		max = myMax(max, dp[i])
	}
	return max
}

func myMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestMaxSumNums(t *testing.T) {
	nums := []int{1, 4, -10, -2, 2, 4}

	res := maxSumNums(nums)

	fmt.Println(res)
}
