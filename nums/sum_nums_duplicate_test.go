package nums

import (
	"fmt"
	"testing"
)

// 参考combine_test.go
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的 组合。
// candidates 中的数字可以无限制重复被选取。(注意不能有0和负数，因为0的存在，会导致有无穷个组合)

// 题目要求出总和为 sum 的所有组合，组合需要去重。 这一题和第 47 题类似，只不过元素可以反复使用。
// 组合求和函数
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int
	backtrack(candidates, target, 0, current, &result)
	return result
}

// 回溯函数  可以逐步递推，看看结果，精髓是，层层深入，然后逐层退出
func backtrack(candidates []int, target int, start int, current []int, result *[][]int) {
	// 当 target 为 0 时，找到一个组合
	if target == 0 {
		// 将当前组合加入结果，防止后续修改current时，导致result结果发生变化
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	// 遍历候选元素
	for i := start; i < len(candidates); i++ {
		// 如果当前候选数大于剩余 target，则剪枝
		if candidates[i] > target {
			continue
		}

		// 将当前候选数加入当前组合
		current = append(current, candidates[i])
		// 递归调用回溯函数，target 减去当前数，且允许重复使用当前数
		backtrack(candidates, target-candidates[i], i, current, result)
		// 回溯，移除最后加入的数，不包含最后一个元素的取法，前闭后开原则
		current = current[:len(current)-1]
	}
}

func combinationSum1(nums []int, target int) [][]int {
	var res [][]int
	var cur []int
	backtrack1(nums, target, 0, &res, cur)
	return res
}

func backtrack1(nums []int, target int, start int, res *[][]int, cur []int) {
	if target == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		if nums[i] > target {
			continue
		}
		cur = append(cur, nums[i])
		backtrack1(nums, target-nums[i], i, res, cur)
		cur = cur[:len(cur)-1]
	}
}

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}

	ans := combinationSum1(candidates, 7)
	fmt.Println(ans)

}
