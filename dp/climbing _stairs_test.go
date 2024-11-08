package dp

import (
	"fmt"
	"testing"
)

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼 顶呢?注意:给定 n 是一个正整数
// 简单的 DP，经典的爬楼梯问题。一个楼梯可以由 n-1 和 n-2 的楼梯爬上来。 这一题求解的值就是斐波那契数列。
func climbStairs(n int) int {
	dp := make([]int, n+1)
	// dp[0], dp[1] = 1, 1：表示如果楼梯为 0 阶（即起点），只有 1 种方法就是不爬。而爬 1 阶楼梯，也只有 1 种方法。
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		//  从第 2 阶开始，使用递推公式 dp[i] = dp[i-1] + dp[i-2]，意思是到达第 i 阶楼梯有两种方式：
		//	从 i-1 阶跳 1 阶上来。
		//	从 i-2 阶跳 2 阶上来。
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(2))
}
