package part1

import (
	"2023/day5/part1/parsers"
	"2023/day5/part1/utils"
)

func Run(input []string) int {
	seeds, categoryMappings := parsers.Parse(input)
	var results []int

	for _, seed := range seeds {
		results = append(results, seed.FindLocation(categoryMappings))
	}

	return utils.FindMinInSlice(results)
}
