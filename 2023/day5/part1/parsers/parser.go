package parsers

import (
	"2023/day5/part1/models"
	"strconv"
	"strings"
)

func Parse(lines []string) ([]models.Seed, []models.RangeMapping) {
	var seeds []models.Seed
	var mappings []models.RangeMapping

	blocks := make(map[int][]string)
	var blockCount int

	for _, line := range lines {
		if line == "" {
			blockCount++
			continue
		}

		blocks[blockCount] = append(blocks[blockCount], line)
	}

	for _, line := range blocks {
		if strings.HasPrefix(line[0], "seeds") {
			seeds = append(seeds, parseSeeds(line[0])...)
		}

		if strings.HasSuffix(line[0], " map:") {
			mappings = append(mappings, parseRangeMappings(line)...)
		}
	}

	return seeds, mappings
}

func toInt(s string) int {
	result, _ := strconv.Atoi(s)

	return result
}
