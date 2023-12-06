package parsers

import (
	"2023/day6/part2/models"
	"2023/day6/part2/utils"
)

func Parse(lines []string) models.Record {
	time := utils.ToInt(utils.RemoveNonDigits(lines[0]))
	distance := utils.ToInt(utils.RemoveNonDigits(lines[1]))

	return models.Record{Time: time, Distance: distance}
}
