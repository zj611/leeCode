package dp

import (
	"fmt"
	"testing"
)

// 最长公共子字符串
func longestCommonSubString(str1, str2 string) string {
	chs1 := len(str1)
	chs2 := len(str2)

	maxLength := 0 //记录最大长度
	end := 0       //记录最大长度的结尾位置

	rows := 0
	cols := chs2 - 1
	for rows < chs1 {
		i, j := rows, cols
		length := 0 //记录长度
		for i < chs1 && j < chs2 {
			if str1[i] != str2[j] {
				length = 0
			} else {
				length++
			}
			if length > maxLength {
				end = i
				maxLength = length
			}
			i++
			j++
		}
		if cols > 0 {
			cols--
		} else {
			rows++
		}
	}
	return str1[(end - maxLength + 1):(end + 1)]
}

func TestLongestCommonSubString(t *testing.T) {
	str1 := "abchsj"
	str2 := "achhjmmmqwa"

	resString := longestCommonSubString(str1, str2)
	fmt.Println(resString)
}
