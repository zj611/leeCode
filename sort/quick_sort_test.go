package sort

import (
	"fmt"
	"testing"
)

func quickSort(nums []int) []int {
	_quickSort1(nums, 0, len(nums)-1)
	return nums
}

func _quickSort1(nums []int, left, right int) {
	if left > right {
		return
	}
	i, j, base := left, right, nums[right]

	for i < j {
		// 以base为分割线，nums[i] 都比它小
		for nums[i] <= base && i < j {
			i++
		}
		// 以base为分割线，nums[j] 都比它大
		for nums[j] >= base && i < j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 保证nums[j]的左边都比它小，右边都比它大，然后再把nums[j]换到最右边
	nums[j], nums[right] = nums[right], nums[j]

	_quickSort1(nums, left, i-1)
	_quickSort1(nums, i+1, right)
}

func TestQuickSort(t *testing.T) {
	var a = []int{11, 7, 3, 4, 0, 1, 8, 0, 6, 5, 6, 7, 8}
	fmt.Println(a)
	b := quickSort(a)
	fmt.Println(b)
}
