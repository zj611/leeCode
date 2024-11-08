package nums

import (
	"fmt"
	"testing"
)

//给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素最多暴露 2 个。最后返回去重以后 数组的⻓度值。

//问题提示有序数组，一般最容易想到使用双指针的解法，双指针的关键点:移动两个指针的条件。
//在该题中移动的条件:快指针从头遍历数组，慢指针指向修改后的数组的末端，当慢指针指向倒数第二个数与
//快指针指向的数不相等时，才移动慢指针，同时赋值慢指针。
//处理边界条件:当数组小于两个元素时，不做处理。

func removeDuplicates2(nums []int) int {
	slow := 0
	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 1, 1, 1, 2}
	fmt.Println(nums)
	ans := removeDuplicates2(nums)
	fmt.Println(ans)

}
