package goroutine

import (
	"fmt"
	"testing"
)

func is_palindrome(str string) bool {
	ch := make(chan byte)
	//defer close(ch)

	length := len(str)
	t := 0
	go func() {
		for i := 0; i < length; i++ {
			t = i
			ch <- str[i]
		}
	}()

	res := true
	for i := length - 1; i >= 0; i-- {
		tt := <-ch
		if tt != str[i] {
			res = false
			break
		}
	}

	// 释放掉通道内数据
	if !res {
		for d := t; d < len(str); d++ {
			<-ch
		}
	}
	close(ch)

	return res
}

func TestIsPalindromee(t *testing.T) {
	fmt.Println(is_palindrome("he9eh"))
}
