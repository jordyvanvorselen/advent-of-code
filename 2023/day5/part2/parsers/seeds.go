package parsers

import (
	"2023/day5/part2/models"
	"2023/day5/part2/utils"
	"strings"
)

func parseSeeds(line string) []models.Seeds {
	var seeds []models.Seeds

	seedNumbers := strings.ReplaceAll(line, "seeds: ", "")
	numbers := strings.Fields(seedNumbers)

	for i, _ := range numbers {
		if i%2 != 0 {
			continue
		}

		start := utils.ToInt(numbers[i])
		length := utils.ToInt(numbers[i+1])

		seeds = append(seeds, models.Seeds{RangeStart: start, RangeLength: length})
	}

	return seeds
}
