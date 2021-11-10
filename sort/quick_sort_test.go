package sort

import (
	"fmt"
	"testing"
)

func quickSort(nums []int) []int {
	_quickSort1(nums, 0, len(nums)-1)
	return nums
}

//func _quickSort(nums []int, left, right int) {
//
//	if left > right {
//		return
//	}
//
//	i, j, base := left, right, nums[right]
//
//	for i < j {
//		for nums[i] <= base && i < j {
//			i++
//		}
//		for nums[j] >= base && i < j {
//			j--
//		}
//
//
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//	nums[j], nums[right] = nums[right], nums[j]
//
//	_quickSort(nums, left, i-1)
//	_quickSort(nums, i+1, right)
//}

func _quickSort1(nums []int, left, right int) {
	if left > right {
		return
	}

	//i, j, base := left, right, nums[left]
	i, j, base := left, right, nums[right]

	for i < j {
		for nums[i] <= base && i < j {
			i++
		}
		for nums[j] >= base && i < j {
			j--
		}
		//for nums[j] >= base && i < j {
		//	j--
		//}
		//for nums[i] <= base && i < j {
		//	i++
		//}
		nums[i], nums[j] = nums[j], nums[i]
	}
	//nums[left], nums[i] = nums[i], nums[left]
	nums[j], nums[right] = nums[right], nums[j]

	_quickSort1(nums, left, i-1)
	_quickSort1(nums, i+1, right)

}

//func quickSort(nums []int) []int {
//
//	_quickSort(nums, 0, len(nums)-1)
//	return nums
//}
//
//func _quickSort(nums []int, left, right int)  {
//	if left > right{
//		return
//	}
//	i,j,base := left,right,nums[left]
//
//	for i < j{
//		for nums[j] >= base && i < j{
//			j--
//		}
//		for nums[i] <= base && i < j{
//			i++
//		}
//		nums[i],nums[j] = nums[j],nums[i]
//	}
//	nums[i], nums[left] = nums[left],nums[i]
//
//	_quickSort(nums,left, i-1)
//	_quickSort(nums,i+1,right)
//
//}

func TestQuickSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 8, 0, 6, 5, 6, 7, 8}
	fmt.Println(a)
	b := quickSort(a)
	fmt.Println(b)
	//var b = []int{222,333,444,555}

}
