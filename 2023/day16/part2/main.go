package part2

import (
	"2023/day16/part2/models"
	"2023/day16/part2/parsers"
	"github.com/samber/lo"
)

func Run(input []string) int {
	var result int

	grid := parsers.Parse(input)
	starts := grid.StartingPositions()

	for _, start := range starts {
		energized := getResult(start, grid)

		if energized > result {
			result = energized
		}
	}

	return result
}

func getResult(start models.Step, grid models.Grid) int {
	var marked []models.Tile
	var cache []models.Step

	steps := []models.Step{start}

	for {
		var next []models.Step
		for _, step := range steps {
			if lo.Contains(cache, step) {
				continue
			}

			cache = append(cache, step)
			next = append(next, grid.GetNext(step)...)
		}

		if len(next) == 0 {
			break
		}

		for _, step := range next {
			if !lo.Contains(marked, step.Tile) {
				marked = append(marked, step.Tile)
			}
		}

		steps = next
	}

	return len(marked)
}
