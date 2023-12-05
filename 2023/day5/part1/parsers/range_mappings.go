package parsers

import (
	"2023/day5/part1/models"
	"strings"
)

func parseRangeMappings(block []string) []models.RangeMapping {
	var rangeMappings []models.RangeMapping

	for _, line := range block[1:] {
		rangeNumbers := strings.Fields(line)

		rangeMappings = append(rangeMappings, models.RangeMapping{
			DestinationRangeStart: toInt(rangeNumbers[0]),
			SourceRangeStart:      toInt(rangeNumbers[1]),
			RangeLength:           toInt(rangeNumbers[2]),
		})
	}

	return rangeMappings
}
