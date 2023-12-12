package parsers

import (
	"2023/day12/part1/models"
	"2023/day12/part1/utils"
	"github.com/samber/lo"
	"strings"
)

func Parse(lines []string) []models.ConditionRecord {
	var result []models.ConditionRecord

	for _, line := range lines {
		inputs := strings.Split(line, " ")

		raw := strings.Split(inputs[0], "")
		broken := lo.Map(strings.Split(inputs[1], ","), func(broken string, index int) models.Group {
			return models.Group(utils.ToInt(broken))
		})

		result = append(result, models.ConditionRecord{
			Raw:    raw,
			Broken: broken,
		})
	}

	return result
}
