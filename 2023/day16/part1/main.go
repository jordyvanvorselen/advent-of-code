package part1

import (
	"2023/day16/part1/models"
	"2023/day16/part1/parsers"
	"github.com/samber/lo"
)

func Run(input []string) int {
	grid := parsers.Parse(input)

	start := models.Step{Tile: grid["0,0"], Direction: models.DOWN}

	marked := []models.Tile{start.Tile}

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
