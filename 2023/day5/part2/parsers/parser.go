package parsers

import (
	"2023/day5/part2/models"
	"strings"
)

func Parse(lines []string) ([]models.Seeds, []models.CategoryMapping) {
	var seeds []models.Seeds
	var mappings []models.CategoryMapping

	blocks := make(map[int][]string)
	var blockCount int

	for _, line := range lines {
		if line == "" {
			blockCount++
			continue
		}

		blocks[blockCount] = append(blocks[blockCount], line)
	}

	for i := 0; i < len(blocks); i++ {
		block := blocks[i]

		if strings.HasPrefix(block[0], "seeds") {
			seeds = append(seeds, parseSeeds(block[0])...)
		}

		if strings.HasSuffix(block[0], " map:") {
			rangeMappings := parseRangeMappings(block)
			from, to := ParseCategories(block[0])
			mappings = append(mappings, models.CategoryMapping{From: from, To: to, Ranges: rangeMappings})
		}
	}

	return seeds, mappings
}
