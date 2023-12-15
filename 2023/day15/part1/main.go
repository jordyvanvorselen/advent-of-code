package part1

import (
	"2023/day15/part1/parsers"
)

func Run(input []string) int {
	var result int

	items := parsers.Parse(input)

	for _, item := range items {
		result += item.Hash()
	}

	return result
}
