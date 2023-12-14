package part2

import (
	"2023/day14/part2/parsers"
)

func Run(input []string) int {
	var result int

	list := parsers.Parse(input)

	for i := 0; i < 1000; i++ {
		newList := list.Cycle()

		list = newList
	}

	for _, col := range list.Columns {
		for i, c := range col {
			if string(c) == "O" {
				result += len(col) - i
			}
		}
	}

	return result
}
