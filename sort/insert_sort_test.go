package sort

import (
	"fmt"
	"testing"
)

func insertSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j-1] > nums[j] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

func TestInsertSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 5, 6, 7, 8}
	fmt.Println(a)
	b := insertSort(a)
	fmt.Println(b)
	//var b = []int{222,333,444,555}

}
