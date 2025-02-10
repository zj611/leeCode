package nums

import (
	"fmt"
	"testing"
)

// 给定一个未排序的整数数组，找出最⻓连续序列的⻓度。要求算法的时间复杂度为 O(n)。

func longestConsecutive(nums []int) int {
	res, numMap := 0, map[int]int{}
	for _, num := range nums {
		if numMap[num] == 0 {
			left, right, sum := 0, 0, 0
			//•	left：num-1 在哈希表中的长度（即以 num-1 为结束的序列的长度）。
			if numMap[num-1] > 0 {
				left = numMap[num-1]
			} else {
				left = 0
			}
			//•	right：num+1 在哈希表中的长度（即以 num+1 为开始的序列的长度）。
			if numMap[num+1] > 0 {
				right = numMap[num+1]
			} else {
				right = 0
			}
			//•	如果 num-1 或 num+1 不存在于哈希表中，其长度默认为 0。

			// sum: length of the sequence n is in
			sum = left + right + 1
			numMap[num] = sum
			// keep track of the max length
			res = max(res, sum)

			// extend the length to the boundary(s) of the sequence
			// will do nothing if n has no neighbors
			numMap[num-left] = sum
			numMap[num+right] = sum
		} else {
			// 重复数字，跳过
			continue
		}
	}
	return res
}

//遍历 nums = [100, 4, 200, 1, 3, 2, 5, 0, 7, 6] 的过程如下：

//1.	num = 100：
//•	无邻居，sum = 1。
//•	更新 numMap = {100: 1}，res = 1。

//2.	num = 4：
//•	无邻居，sum = 1。
//•	更新 numMap = {100: 1, 4: 1}，res = 1。

//3.	num = 200：
//•	无邻居，sum = 1。
//•	更新 numMap = {100: 1, 4: 1, 200: 1}，res = 1。

//4.	num = 1：
//•	无邻居，sum = 1。
//•	更新 numMap = {100: 1, 4: 1, 200: 1, 1: 1}，res = 1。

//5.	num = 3：
//•	右邻居 4 存在，sum = 1 + 1 = 2。
//•	更新 numMap = {100: 1, 4: 2, 200: 1, 1: 1, 3: 2}，res = 2。

//6.	num = 2：
//•	左邻居 1 和右邻居 3 存在，sum = 1 + 2 + 1 = 4。
//•	更新 numMap = {100: 1, 4: 4, 200: 1, 1: 4, 3: 4, 2: 4}，res = 4。

//7.	num = 5：
//•	左邻居 4 存在，sum = 4 + 1 = 5。
//•	更新 numMap = {100: 1, 4: 5, 200: 1, 1: 5, 3: 4, 2: 5, 5: 5}，res = 5。

//8.	num = 0：
//•	右邻居 1 存在，sum = 5 + 1 = 6。
//•	更新 numMap = {100: 1, 4: 5, 200: 1, 1: 6, 3: 4, 2: 5, 5: 5, 0: 6}，res = 6。

//9.	num = 7：
//•	无邻居，sum = 1。
//•	更新 numMap = {100: 1, 4: 5, 200: 1, 1: 6, 3: 4, 2: 5, 5: 5, 0: 6, 7: 1}，res = 6。

//10.	num = 6：
//•	左邻居 5 和右邻居 7 存在，sum = 5 + 1 + 1 = 8。
//•	更新 numMap = {100: 1, 4: 8, 200: 1, 1: 8, 3: 4, 2: 8, 5: 8, 0: 8, 7: 8, 6: 8}，res = 8。

func TestLongestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2, 5, 0, 7, 6}))
}
