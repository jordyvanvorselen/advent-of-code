package parsers

import (
	"2023/day5/part1/models"
	"2023/day5/part1/utils"
	"strings"
)

func parseRangeMappings(block []string) []models.RangeMapping {
	var rangeMappings []models.RangeMapping

	for _, line := range block[1:] {
		rangeNumbers := strings.Fields(line)

		rangeMappings = append(rangeMappings, models.RangeMapping{
			DestinationRangeStart: utils.ToInt(rangeNumbers[0]),
			SourceRangeStart:      utils.ToInt(rangeNumbers[1]),
			RangeLength:           utils.ToInt(rangeNumbers[2]),
		})
	}

	return rangeMappings
}
