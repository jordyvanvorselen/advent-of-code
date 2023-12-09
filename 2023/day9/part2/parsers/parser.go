package parsers

import (
	"2023/day5/part2/utils"
	"2023/day9/part2/models"
	"github.com/samber/lo"
	"strings"
)

func Parse(lines []string) []models.Sequence {
	var result []models.Sequence

	for _, line := range lines {
		sequence := lo.Map(strings.Fields(line), func(item string, index int) int {
			return utils.ToInt(item)
		})

		result = append(result, sequence)
	}

	return result
}
