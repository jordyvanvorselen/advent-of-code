package part1

import (
	"2023/day12/part1/parsers"
)

func Run(input []string) int {
	var result int

	records := parsers.Parse(input)

	for _, record := range records {
		result += record.PossibleBrokenCombinations()
	}

	return result
}
