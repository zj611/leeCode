package nums

import (
	"fmt"
	"testing"
)

// 给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素只有一个。最后返回去重以后数组 的⻓度值。
func removeDuplicates(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nums
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1, nums
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1, nums
}

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		// 当不满足这个条件时，j不走了，等着后面的元素来替换
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	fmt.Println(nums)
	return j
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{-1, 0, 1, 1, 0, 2}
	fmt.Println(nums)
	ans, nums1 := removeDuplicates(nums)
	fmt.Println(ans, nums1)

	nums2 := []int{-1, 0, 1, 1, 0, 2}
	removeElement(nums2, 1)
}
