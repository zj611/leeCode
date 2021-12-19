package dp

import (
	"fmt"
	"testing"
)

//func longestCommonSubsequence(text1 string, text2 string) (int,string) {
//	var resSubString string
//	m,n := len(text1), len(text2)
//
//	dp := make([][]int, m +1)
//
//	for i := range dp{
//		dp[i] = make([]int, n+1)
//	}
//
//	for k1,v1 := range text1{
//		for k2,v2 := range text2{
//			if v1 == v2 {
//				resSubString += string(v1)
//				dp[k1+1][k2+1] = dp[k1][k2] + 1
//			}else{
//				dp[k1+1][k2+1] = max(dp[k1][k2+1], dp[k1+1][k2])
//			}
//		}
//	}
//	return dp[m][n],resSubString
//}

func longestCommonSubsequence(text1 string, text2 string) (int, string) {

	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			if text1[x-1] == text2[y-1] {
				dp[x][y] = dp[x-1][y-1] + 1
			} else {
				dp[x][y] = max(dp[x-1][y], dp[x][y-1])
			}
		}
	}

	//搜索公共子序列的时候，从右下角开始，向上、左上和左三个方向移动，
	//如果当前位置的值跟左边一个或者上边一个值相同，则向左或者向上移动，
	//如果不等，则向左上移动，它们的差为1，搜索代码如下：
	resCount := dp[m][n]

	resSubString := make([]byte, resCount)
	index := len(resSubString) - 1
	for index >= 0 {
		if n > 0 && dp[m][n] == dp[m][n-1] {
			n--
		} else if m > 0 && dp[m][n] == dp[m-1][n] {
			m--
		} else {
			resSubString[index] = byte(text1[m-1])
			index--
			m--
			n--
		}
	}

	//return dp[m][n], string(resSubString)
	return resCount, string(resSubString)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestLongestCommonSubSequence(t *testing.T) {
	str1 := "abchhsj"
	str2 := "achhjmmmqwa"

	resNum, resString := longestCommonSubsequence(str1, str2)

	fmt.Println(resNum, resString)

	//golang中string底层是通过byte数组实现的。
	//中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
	fmt.Println("----------")
	s1 := "sdsd"
	s2 := "sd长进"
	fmt.Println(len(s1), " ", len(s2)) // 4 8

	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	s3 := []rune(s2)
	fmt.Println(len(s3)) //4

}
