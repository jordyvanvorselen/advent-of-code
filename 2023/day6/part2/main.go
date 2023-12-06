package part2

import (
	"2023/day6/part2/models"
	"2023/day6/part2/parsers"
)

func Run(input []string) int {
	record := parsers.Parse(input)
	boat := models.Boat{StartingSpeed: 0, HoldingIncrease: 1}

	return record.WaysToBeatUsing(boat)
}
