package algorithms

import "math"

func BinarySearch(needle int, haystack []int) bool {
	l := 0
	r := len(haystack) - 1
	for l < r {
		m := l + int(math.Floor((float64(r)-float64(l))/2.0))
		if needle == haystack[m] {
			return true
		} else if needle > haystack[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return false
}
