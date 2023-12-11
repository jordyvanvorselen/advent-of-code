package part1

import (
	"2023/day11/part1/parsers"
	"fmt"
)

func Run(input []string) int {
	var result int

	graph, galaxies := parsers.Parse(input)

	pairs := findUniquePairs(galaxies)

	fmt.Println(pairs)

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
