package part1

import (
	"2023/day14/part1/models"
	"2023/day14/part1/parsers"
)

func Run(input []string) int {
	var result int

	lists := parsers.Parse(input)

	for _, list := range lists {
		newList := list.Tilt(models.NORTH)

		for _, col := range newList.Columns {
			for i, c := range col {
				if string(c) == "O" {
					result += len(col) - i
				}
			}
		}
	}

	return result
}
