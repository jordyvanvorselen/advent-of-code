package part2

import "2023/day9/part2/parsers"

func Run(input []string) int {
	var result int

	sequences := parsers.Parse(input)

	for _, sequence := range sequences {
		result += sequence.Prediction()
	}

	return result
}
