package part1

import "cmp"

func BOverlapCountInA(a, b []int) int {
	var result int

	for _, e := range b {
		if elementInSlice(e, a) {
			result++
		}
	}

	return result
}

func elementInSlice[T cmp.Ordered](a T, list []T) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
