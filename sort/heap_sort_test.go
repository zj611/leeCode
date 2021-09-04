package sort

import (
	"fmt"
	"testing"
)

func heapSort(nums []int) []int {
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		sink(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		sink(nums, 0, end)
	}
	return nums
}

func sink(heap []int, root, end int) {
	for {
		child := root*2 + 1
		if child > end {
			return
		}
		if child < end && heap[child] <= heap[child+1] {
			child++
		}
		if heap[root] > heap[child] {
			return
		}
		heap[root], heap[child] = heap[child], heap[root]
		root = child
	}
}

func heapSort1(nums []int) []int {
	n := len(nums)
	var i int
	for i = (n - 1) / 2; i >= 0; i-- {
		shiftDown(nums, n, i)
	}

	for i = n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		shiftDown(nums, i, 0)
	}
	return nums
}

func shiftDown(nums []int, n, k int) {
	var j int
	for 2*k+1 < n {
		j = 2*k + 1
		if j+1 < n && nums[j+1] > nums[j] {
			j++
		}
		if nums[k] >= nums[j] {
			break
		}
		nums[k], nums[j] = nums[j], nums[k]
		k = j
	}
}

func TestHeapSort(t *testing.T) {
	var a = []int{3, 4, 0, 1, 8, 0, 6, 5, 6, 7, 8}
	fmt.Println(a)
	b := heapSort(a)
	fmt.Println(b)
	//var b = []int{222,333,444,555}

}
