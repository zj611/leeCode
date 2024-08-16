package binary_search

import (
	"fmt"
	"testing"
)

//二分查找

func TestBinarySearch(t *testing.T) {
	//re := guessNumber(80)
	re := guessNumber1(80)
	fmt.Println(re)
}

// 给点一个数n，查找其中一个目标数m， m在1～n之间
func guessNumber1(n int) int {

	if guess(n) == 0 {
		return n
	}

	l := 1
	r := n
	mid := 0

	for l < r {
		mid = l + (r-l)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r
}

func guess1(n int) int {
	if guess(n) == 0 {
		return n
	}
	l := 1
	r := n
	for l < r {
		mid := l + (r-l)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return r
}

const (
	target = 10 // 目标查找数
)

func guess(a int) int {
	if a == target {
		return 0
	} else if a > target {
		return -1
	} else {
		return 1
	}
}
