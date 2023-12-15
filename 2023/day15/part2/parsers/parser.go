package parsers

import (
	"2023/day15/part2/models"
	"2023/day15/part2/utils"
	"strings"
)

func Parse(lines []string) []models.InitItem {
	var result []models.InitItem

	for _, line := range strings.Split(lines[0], ",") {
		if strings.Contains(line, "-") {
			result = append(result, models.InitItem{
				Label: line[:len(line)-1],
				Op:    models.REMOVE,
			})

			continue
		}

		parts := strings.Split(line, "=")
		result = append(result, models.InitItem{
			Label:       parts[0],
			Op:          models.ADD,
			FocalLength: utils.ToInt(parts[1]),
		})
	}

	return result
}
