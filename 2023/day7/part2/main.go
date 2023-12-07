package part2

import (
	"2023/day7/part2/parsers"
	"sort"
)

func Run(input []string) int {
	var result int

	hands := parsers.Parse(input)
	sort.Sort(hands)

	for i, hand := range hands {
		result += hand.Bid * (i + 1)
	}

	return result
}
