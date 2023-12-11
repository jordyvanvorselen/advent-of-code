package part2

import (
	"2023/day11/part2/parsers"
)

func Run(input []string, factor int) int {
	var result int

	graph, galaxies := parsers.Parse(input, factor)

	pairs := findUniquePairs(galaxies)

	for _, pair := range pairs {
		a := pair[0]
		b := pair[1]

		if a == b {
			continue
		}

		best, err := graph.Shortest(a, b)

		if err != nil {
			panic(err)
		}

		result += int(best.Distance)
	}

	return result
}

func findUniquePairs(slice []int) [][]int {
	var uniquePairs [][]int

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := []int{slice[i], slice[j]}
			uniquePairs = append(uniquePairs, pair)
		}
	}

	return uniquePairs
}
