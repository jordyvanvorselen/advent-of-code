package utils

import "strconv"

func ToInt(s string) int {
	result, _ := strconv.Atoi(s)

	return result
}
