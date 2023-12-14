package part1

import (
	"2023/day14/part1/parsers"
	"fmt"
)

func Run(input []string) int {
	var result int

	lists := parsers.Parse(input)

	for _, list := range lists {
		newList := list.TiltNorth()

		fmt.Println(newList)
	}

	return result
}
