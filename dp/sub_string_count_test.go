package dp

import (
	"fmt"
	"testing"
)

// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列是指，通过删除一些 (也可以不删除)字符且不干扰剩余字符相对位置所组成的新字符串。
// (例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是)题目数据保证答案符合 32 位带符号整数范围。

// 解法一 压缩版 DP
func subStringCount(s string, t string) int {
	dp := make([]int, len(s)+1) // 初始化 DP 数组，大小为 len(s) + 1，默认值为 0

	// 遍历 t 的每个字符
	for i, curT := range t {
		pre := 0                 // 记录上一次 dp[j+1] 的值，用于当前计算
		for j, curS := range s { // 遍历 s 的每个字符
			if i == 0 { // 如果是 t 的第一个字符，子问题可以初始化
				pre = 1
			}
			newDP := dp[j+1]  // 暂存 dp[j+1]，因为当前 dp[j+1] 会被更新
			if curT == curS { // 如果 t[i] 和 s[j] 匹配
				dp[j+1] = dp[j] + pre
			} else { // 如果不匹配，沿用 dp[j] 的值
				dp[j+1] = dp[j]
			}
			pre = newDP // 更新 pre，用于下一个 j
		}
	}

	return dp[len(s)] // 返回最后一个 dp 值
}

// 解法二 普通 DP
// 动态规划方法：
//  1. 定义 dp[i][j]：
//     表示 t[j:]（t 的后缀）在 s[i:]（s 的后缀）中出现的次数。
//  2. 初始化：
//     当 t 是空字符串时（j == n），它是 s 的子序列，只有一种情况：dp[i][n] = 1。
//     当 s 是空字符串时（i == m），t 无法作为 s 的子序列，dp[m][j] = 0（j < n）。
//  3. 转移方程：
//     如果 s[i] == t[j]，dp[i][j] 的值由两部分构成：
//     匹配当前字符，继续比较后续部分：dp[i+1][j+1]。
//     跳过当前字符，只看 s[i+1:] 部分：dp[i+1][j]。
//     所以，dp[i][j] = dp[i+1][j+1] + dp[i+1][j]。
//     如果 s[i] != t[j]，只能跳过 s[i]，dp[i][j] = dp[i+1][j]。
func numDistinct1(s, t string) int {
	ss, tt := len(s), len(t)
	if ss < tt {
		return 0
	}
	dp := make([][]int, ss+1)

	for i := range dp {
		dp[i] = make([]int, tt+1)
		//dp[i][j]表示 t[j:]（t 的后缀）在 s[i:]（s 的后缀）中出现的次数。
		dp[i][tt] = 1
	}
	// 倒序往前，这是重点
	for i := ss - 1; i >= 0; i-- { // m-s
		for j := tt - 1; j >= 0; j-- { // n-t
			if s[i] == t[j] {
				// 相等的话，注意这里已经到第i层了，所以i+1层一定有值
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

func TestSubStringCount(t *testing.T) {
	fmt.Println(subStringCount("ABCCCDE", "ACE"))
}
