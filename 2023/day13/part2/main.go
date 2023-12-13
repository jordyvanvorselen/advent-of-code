package part2

import (
	"2023/day13/part2/parsers"
)

func Run(input []string) int {
	var result int

	mirrorLists := parsers.Parse(input)

	for _, ml := range mirrorLists {
		indexV, foundV, isOrgV := ml.FindPerfectColumnReflectionIndex()
		indexH, foundH, isOrgH := ml.FindPerfectRowReflectionIndex()

		if foundV && !isOrgV {
			result += indexV + 1
		} else if foundH && !isOrgH {
			result += (indexH + 1) * 100
		} else if foundV && isOrgV {
			result += indexV + 1
		} else if foundH && isOrgH {
			result += (indexH + 1) * 100
		}
	}

	return result
}
