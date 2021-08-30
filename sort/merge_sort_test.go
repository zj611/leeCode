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
	//把left和right剩下的元素拷到result
	copy(result[i:], left[l:])
	copy(result[(i+len(left)-l):], right[r:])
	return result

}

//func mergeSort(nums []int) {
//	_mergeSort(nums, 0, len(nums)-1)
//}

//func _mergeSort(nums []int, l, r int) {
//	if l >= r {
//		return
//	}
//
//	mid := (l + r) / 2
//	_mergeSort(nums, l, mid)
//	_mergeSort(nums, mid+1, r)
//	merge(nums, l, mid, r)
//}
//
//func merge(nums []int, l, mid, r int) {
//	var tmparr []int
//	var s1, s2 = l, mid + 1
//	for s1 <= mid && s2 <= r {
//		if nums[s1] > nums[s2] {
//			tmparr = append(tmparr, nums[s2])
//			s2++
//		} else {
//			tmparr = append(tmparr, nums[s1])
//			s1++
//		}
//	}
//	if s1 <= mid {
//		tmparr = append(tmparr, nums[s1:mid+1]...)
//	}
//	if s2 <= r {
//		tmparr = append(tmparr, nums[s2:r+1]...)
//	}
//	for pos, item := range tmparr {
//		nums[l+pos] = item
//	}
//}

func TestMergeSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 5, 6, 7, 8}
	b := mergeSort(a)
	fmt.Println(b)
	//var b = []int{222,333,444,555}

}
