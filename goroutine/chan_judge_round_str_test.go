package goroutine

import (
	"fmt"
	"testing"
)

func is_palindrome(str string) bool {
	ch := make(chan byte)
	//defer close(ch)

	length := len(str)

	go func() {
		for i := 0; i < length; i++ {
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

	return res
}

func TestIsPalindromee(t *testing.T) {
	fmt.Println(is_palindrome("he9eh"))
}
