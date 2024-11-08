package dp

import (
	"fmt"
	"testing"
)

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组(子数组最少包含一个元素)，返回其最大和
//题目要求输出数组中某个区间内数字之和最大的那个值。 dp[i] 表示 [0,i] 区间内各个子区间和的最大值，
//状态转移方程是 dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0) ， dp[i] = nums[i] (dp[i-1] ≤ 0)。

// 解法一 DP
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := 0

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = max(dp[i], res)

	}

	return res
}

func TestMaxSubArray(t *testing.T) {
	res := maxSubArray1([]int{1, 3, -10, 6, 7})
	fmt.Println(res)
}
