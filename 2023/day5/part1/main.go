package part1

import (
	"2023/day5/part1/parsers"
)

func Run(input []string) int {
	var result int

	seeds, rangeMappings := parsers.Parse(input)

	for _, seed := range seeds {
		location := seed.FindLocation()
	}

	return result
}
