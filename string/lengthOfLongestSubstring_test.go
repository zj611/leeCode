package string

import (
	"fmt"
	"testing"
)

// 给定一个字符串，寻找没有重复字母的最⻓子串
// 解法三 滑动窗口-哈希桶
func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	l, r := 0, 0
	indexes := make(map[byte]int, len(s))
	for right < len(s) {
		// 如果在 right节点往前走的时候，存在已经走过的点，并且id还比left大，则left调整为id+x
		// 这样可以保证left和right之间没有重复点
		if id, ok := indexes[s[right]]; ok && id >= left {
			left = id + 1
		}
		// right节点一直往前走，边走边存
		indexes[s[right]] = right
		right++

		// 记录下最长的间距
		if right-left > res {
			res = right - left
			l = left
			r = right
		}
		//res = max(res, left-right)
	}
	fmt.Println("res", s[l:r])
	return res
}

func lengthOfLongestSubstring2(s string) int {
	indexs := make(map[byte]int, len(s))
	left, right, l, r, res := 0, 0, 0, 0, 0

	for right < len(s) {
		if id, ok := indexs[s[right]]; ok && id >= left {
			left = id + 1
		}
		indexs[s[right]] = right
		right++

		if right-left > res {
			res = right - left
			l = left
			r = right
		}
	}

	println(s[l:r])
	return res
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("sdsdssss9s12133"))
	fmt.Println(lengthOfLongestSubstring2("sdsdssss9s12133"))
}
