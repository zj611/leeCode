package dp

import (
	"fmt"
	"testing"
)

//最大连续子序列和
//func maxSumNums(nums []int) int {
//	n := len(nums)
//
//	if n == 0 {
//		return 0
//	}
//
//	dp := make([]int, n)
//	max := nums[0]
//	dp[0] = nums[0]
//	for i := 1; i < n; i++ {
//		if dp[i-1] > 0 {
//			dp[i] = dp[i-1] + nums[i]
//		} else {
//			dp[i] = nums[i]
//		}
//		max = myMax(max, dp[i])
//	}
//	return max
//}

func maxSumNums(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	max := nums[0]
	dp[0] = nums[0]
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
			r = i
		} else {
			dp[i] = nums[i]
			l = i
		}
		max = myMax(max, dp[i])
	}
	fmt.Println(l, r)
	return max
}

func myMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func subarraySum(nums []int, k int) int {
	hash := map[int]int{}
	hash[0]++
	sum, res := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += hash[sum-k] //如果被命中，说明当前sum和被命中的sum之间差值，等于存在的一个子序列和
		//两个sum之差，说明一定有子序列
		hash[sum]++
		fmt.Println(sum, res, hash)
	}
	return res

}

func TestMaxSumNums(t *testing.T) {
	nums := []int{1, -4, 10, 2, 2, -4}

	res := maxSumNums(nums)

	fmt.Println(res)

	r := subarraySum(nums, res)
	fmt.Println(r)
}
