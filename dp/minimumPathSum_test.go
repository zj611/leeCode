package dp

import (
	"fmt"
	"testing"
)

// 最小路径和 64
// 给定一个包含非负整数的mxn网格，请找出一条从左上⻆到右下⻆的路径，使得路径上的数字总和为最小。说明:每次只能向下或者向右移动一步。

// 解法一 原地 DP，无辅助空间
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 左侧列和上侧行，均累加起来，目的是为了保证当grid[1][1]时，grid[0][1]和，grid[1][0]的值已经是路径和了
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// 开始计算
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

// 一个机器人位于一个 m x n 网格的左上⻆ (起始点在下图中标记为“Start” )。机器人每次只能向下或者向右移动 一步。
// 机器人试图达到网格的右下⻆(在下图中标记为“Finish”)，那么从左上⻆到右下⻆将会有多少条不同的路径?
func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 对于边缘的格子，它只有一种方式到达该位置，即一直竖着走，或一直者横着走
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}

func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(3, 3))
}

// 现在考虑网格中有障碍物。那么从左上⻆到右下⻆将会有多少条不同的路径?
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 处理边界，如果边界上有障碍物，那么后续位置都无法到达
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if dp[0][i-1] != 0 && obstacleGrid[0][i] != 1 {
			dp[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		if dp[i-1][0] != 0 && obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		}
	}

	// 计算
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
