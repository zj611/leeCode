package nums

import (
	"fmt"
	"sort"
	"testing"
)

//给一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，
//使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//注意：答案中不可以包含重复的三元组。

func threeSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] { //当数组索引大于0时，避免重复此轮，跳过
			continue
		}

		third := n - 1
		t := target - nums[first]

		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] { //当数组索引大于0时，避免重复
				continue
			}

			for second < third && nums[second]+nums[third] > t { // 找到一个third就行
				third--
			}

			if second == third { //当third减到和second相同时
				break
			}

			if nums[second]+nums[third] == t {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 1, 0, 2}

	ans := threeSum(nums, 2)
	fmt.Println(ans)

}
