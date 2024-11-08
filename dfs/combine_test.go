package dfs

import (
	"fmt"
	"testing"
)

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。 参考 sum_nums_duplicate_test.go
func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	var res [][]int
	var current []int
	generateCombinations(n, k, 1, current, &res)
	return res
}
func generateCombinations(n, k, start int, current []int, res *[][]int) {
	if len(current) == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, temp)
		return
	}
	// i will at most be n - (k - c.size()) + 1   i最多到这个值
	// 例如 n=5 k=3
	// 如果 current 长度为 1（当前选了一个元素），还需要 2 个元素（k - len(current) = 2）
	// i 最多只能到 n - (k - len(current)) + 1，即 5 - 2 + 1 = 4。
	// 因为当current=1时，i>4还没找到一个合适的，也就是不能满足总长度为3了

	for i := start; i <= n-(k-len(current))+1; i++ { // 剪枝
		current = append(current, i)
		generateCombinations(n, k, i+1, current, res)
		current = current[:len(current)-1]
	}
	return
}

//func TestBinarySearch(t *testing.T) {

func TestCombine(t *testing.T) {
	fmt.Println(combine(4, 2))
}
