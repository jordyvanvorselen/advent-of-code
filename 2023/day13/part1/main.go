package part1

import (
	"2023/day13/part1/parsers"
)

func Run(input []string) int {
	var result int

	mirrorLists := parsers.Parse(input)

	for _, ml := range mirrorLists {
		vertical, foundV := ml.FindPerfectColumnReflectionIndex()
		horizontal, foundH := ml.FindPerfectRowReflectionIndex()

		if foundV {
			result += vertical + 1
		}

		if foundH {
			result += (horizontal + 1) * 100
		}
	}

	return result
}
