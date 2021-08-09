package binary_search

import (
	"fmt"
	"testing"
)

//二分查找

func TestBinarySearch(t *testing.T) {
	re := guessNumber(7)
	fmt.Println(re)
}

func guessNumber(n int) int {

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

const (
	b = 10
)

func guess(a int) int {
	if a == b {
		return 0
	} else if a > b {
		return -1
	} else {
		return 1
	}
}
