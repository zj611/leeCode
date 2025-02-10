package nums

import (
	"fmt"
	"testing"
)

// 【一般题目】合并两个已经有序的数组，结果放在第一个数组中，第一个数组假设空间足够大。要求算法时间复杂度足够低。
// 为了不大量移动元素，就要从2个数组⻓度之和的最后一个位置开始，依次选取两个数组中大的数，
// 从第一个数组 的尾巴开始往头放，只要循环一次以后，就生成了合并以后的数组了。

//Input:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//Output: [1,2,2,3,5,6]

func mergeTwoOrderedNums(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

func TestMergeTwoOrderedNums(t *testing.T) {
	nums1 := []int{1, 2, 4, 6, 7, 0, 0, 0, 0, 0}
	nums2 := []int{1, 3, 9, 16, 17}
	mergeTwoOrderedNums1(nums1, 5, nums2, 5)
	fmt.Println(nums1, nums2)
}

func mergeTwoOrderedNums1(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}

	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}

}
