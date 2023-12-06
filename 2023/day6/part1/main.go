package part1

import (
	"2023/day6/part1/models"
	"2023/day6/part1/parsers"
)

func Run(input []string) int {
	result := 1

	records := parsers.Parse(input)
	boat := models.Boat{StartingSpeed: 0, HoldingIncrease: 1}

	for _, record := range records {
		result = result * record.WaysToBeatUsing(boat)
	}

	return result
}
