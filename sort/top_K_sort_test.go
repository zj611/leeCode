package sort

import (
	"fmt"
	"testing"
)

func topKSort(nums []int, k int) int {
	n := len(nums)
	i, j, base := 0, n-1, nums[n-1]

	for i < j {
		for nums[i] >= base && i < j {
			i++
		}
		for nums[j] <= base && i < j {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if j == k-1 {
		return nums[n-1]
	} else if j < k-1 {
		nums[j], nums[n-1] = nums[n-1], nums[j]
		return topKSort(nums[j+1:], k-(j+1))
	} else {
		return topKSort(nums[:j], k)
	}
}

func TestTopKQuickSort(t *testing.T) {
	var a = []int{0, 1, 3, 2, 4}
	fmt.Println(a)
	b := topKSort(a, 4)
	fmt.Println(b)
	//var b = []int{222,333,444,555}

}
