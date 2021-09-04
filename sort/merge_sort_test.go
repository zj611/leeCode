package sort

import (
	"fmt"
	"testing"
)

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)

}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r, i := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}

	copy(result[i:], left[l:])
	copy(result[i+len(left)-l:], right[r:])
	return result

}

//func mergeSort(nums []int) []int {
//	if len(nums) <= 1 {
//		return nums
//	}
//
//	mid := len(nums) / 2
//	left := mergeSort(nums[:mid])
//	right := mergeSort(nums[mid:])
//
//	return merge(left, right)
//}
//
//func merge(left, right []int) []int {
//	result := make([]int, len(left)+len(right))
//
//	l, r, i := 0, 0, 0
//
//	for l < len(left) && r < len(right) {
//		if left[l] < right[r] {
//			result[i] = left[l]
//			l++
//		} else {
//			result[i] = right[r]
//			r++
//		}
//		i++
//	}
//	//把left和right剩下的元素拷到result
//	copy(result[i:], left[l:])
//	copy(result[(i+len(left)-l):], right[r:])
//	return result
//
//}

func TestMergeSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 5, 6, 4, 7, 6, 8}
	fmt.Println(a)
	b := mergeSort(a)
	fmt.Println(b)
	//var b = []int{222,333,444,555}

}
