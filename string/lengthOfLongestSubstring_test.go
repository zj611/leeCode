package string

import "fmt"

func lengthOfLongestSubstring(s string) int {

	n := len(s)
	if n <= 1 {
		return n
	}

	window := make(map[byte]bool)

	ll, rr := 0, 0
	l, r := 0, 0

	res := 0

	for r < n {
		c := s[r]

		for window[c] {
			delete(window, s[l])
			l++
		}

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
