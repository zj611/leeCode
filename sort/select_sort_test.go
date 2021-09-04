package sort

import (
	"fmt"
	"testing"
)

func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func TestSelectSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 8, 0, 6, 5, 6, 7, 8}
	fmt.Println(a)
	b := selectSort(a)
	fmt.Println(b)
	//var b = []int{222,333,444,555}

}
