package nums

import (
	"fmt"
	"testing"
)

// 二分搜索变种
// 给出一个有序数组 nums 和一个数 target ，要求在数组中找到第一个和这个元素相等的元素下标，最后一个和这个元素相等的元素下标。

// 这一题是经典的二分搜索变种题。二分搜索有 4 大基础变种题:
// 1. 查找第一个值等于给定值的元素
// 2. 查找最后一个值等于给定值的元素
// 3. 查找第一个大于等于给定值的元素
// 4. 查找最后一个小于等于给定值的元素
// 这一题的解题思路可以分别利用变种 1 和变种 2 的解法就可以做出此题。或者用一次变种 1 的方法，
// 然后循环往后找到最后一个与给定值相等的元素。不过后者这种方法可能会使时间复杂度下降到 O(n)，因为有可能 数组中 n 个元素都和给定元素相同。(4 大基础变种的实现⻅代码)
func searchRange(nums []int, target int) []int {
	//return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
	return []int{searchLastLessElement(nums, target), searchFirstGreaterElement(nums, target)}
}

func TestSearchRange(t *testing.T) {
	res := searchRange([]int{1, 3, 4, 5, 5, 6, 7, 8, 9, 10, 22, 32}, 7)
	fmt.Println(res)
}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		//mid := (low + high) / 2
		//这种方法有一个潜在问题：当 low 和 high 都是非常大的整数时，low + high 可能会发生 整型溢出，
		//即结果超过了整数类型的表示范围，导致计算出错误的中间值。
		mid := low + ((high - low) >> 1)
		// 使用 右移运算符 >> 来实现除以 2，这是更高效的操作。 low +：将 low 加回去，得到正确的中点值。
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { // 找到第一个与 target 相等的元素
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的 元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 二分查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {

			if (mid == len(nums)-1) || (nums[mid+1] > target) { // 找到最后一个小于等于 target 的 元素
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
