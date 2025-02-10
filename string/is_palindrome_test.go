package string

import (
	"fmt"
	"strings"
	"testing"
)

// 判断一个字符串是否是回文串，回文串的定义是正着读和反着读都相同，比如 “radar” 或 “level”。
func isPalindrome(s string) bool {
	// 预处理：去掉非字母数字字符并转为小写
	filtered := ""
	for _, char := range s {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
			filtered += strings.ToLower(string(char))
		}
	}

	// 判断是否回文
	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome("level"))                          // true
}
