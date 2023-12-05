package parsers

import (
	"2023/day5/part1/models"
	"strconv"
	"strings"
)

func parseSeeds(line string) []models.Seed {
	var seeds []models.Seed

	seedNumbers := strings.ReplaceAll(line, "seeds: ", "")

	for _, s := range strings.Fields(seedNumbers) {
		number, _ := strconv.Atoi(s)
		seeds = append(seeds, models.Seed{Number: number})
	}

	return seeds
}
