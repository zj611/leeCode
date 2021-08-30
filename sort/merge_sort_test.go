package sort

import (
	"fmt"
	"testing"
)

func mergeSort(nums []int) {
	_mergeSort(nums, 0, len(nums)-1)
}

func _mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	_mergeSort(nums, l, mid)
	_mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	var tmparr = []int{}
	var s1, s2 = l, mid + 1
	for s1 <= mid && s2 <= r {
		if nums[s1] > nums[s2] {
			tmparr = append(tmparr, nums[s2])
			s2++
		} else {
			tmparr = append(tmparr, nums[s1])
			s1++
		}
	}
	if s1 <= mid {
		tmparr = append(tmparr, nums[s1:mid+1]...)
	}
	if s2 <= r {
		tmparr = append(tmparr, nums[s2:r+1]...)
	}
	for pos, item := range tmparr {
		nums[l+pos] = item
	}
}

func TestMergeSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 5, 6, 7, 8}
	mergeSort(a)
	fmt.Println(a)
}
