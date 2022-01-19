package nums

import (
	"fmt"
	"sort"
	"testing"
)

//给一个包含 n 个整数的数组nums，判断nums中是否存在两个元素 a，b，c ，
//使得a + b  = 0 ？请你找出所有和为 0 且不重复的两元组。
//注意：答案中不可以包含重复的两元组。

func twoSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 2 {
		return [][]int{}
	}

	sort.Ints(nums)

	ans := make([][]int, 0)

	second := n - 1

	for first := 0; first < n; first++ {
		if first > 1 && nums[first] == nums[first-1] { //当数组索引大于0时，避免重复
			continue
		}

		for first < second && nums[first]+nums[second] > target { // 找到一个second就行
			second--
		}

		if first == second { //当second减到和first相同时
			break
		}

		if nums[first]+nums[second] == target {
			ans = append(ans, []int{nums[first], nums[second]})
		}
	}

	return ans
}

func TestTwoSum(t *testing.T) {
	nums := []int{-1, 0, 1, 1, 0, 2}

	ans := twoSum(nums, 1)
	fmt.Println(ans)
}
