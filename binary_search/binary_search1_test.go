package binary_search

import "testing"

//二分查找

func TestBinarySearch(t *testing.T) {

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
