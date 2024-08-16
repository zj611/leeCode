package string

import (
	"fmt"
	"testing"
)

// 滑动窗口 找最长不重复子字符串
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	window := make(map[byte]bool)

	l, r := 0, 0

	// 用于记录当前的首末索引位置结果
	ll, rr := 0, 0
	res := 0
	for r < n {
		c := s[r]
		// l递增，确保删除之前全部的重复字符
		for window[c] {
			delete(window, s[l])
			l++
		}
		// 此处确保取最长不重复子字符串
		if res < (r - l + 1) {
			res = r - l + 1
			ll, rr = l, r
		}
		window[c] = true
		r++
	}
	fmt.Println("res", s[ll:rr+1])
	return res
}

func TestLengthOfLongestSubstring(t *testing.T) {
	lengthOfLongestSubstring("sdsdsssss12133")
}
