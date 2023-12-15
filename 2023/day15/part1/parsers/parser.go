package parsers

import (
	"2023/day15/part1/models"
	"strings"
)

func Parse(lines []string) []models.InitItem {
	var result []models.InitItem

	for _, line := range strings.Split(lines[0], ",") {
		result = append(result, models.InitItem{Raw: line})
	}

	return result
}
