package main

import (
	day1part1 "aoc/2023/day1/part1"
	day1part2 "aoc/2023/day1/part2"
	day2part1 "aoc/2023/day2/part1"
	day2part2 "aoc/2023/day2/part2"
	day3part1 "aoc/2023/day3/part1"
	"fmt"
)

func main() {
	fmt.Println("Day 1")
	fmt.Println("Part 1: ", day1part1.Run(readFile("day1/input/data")))
	fmt.Println("Part 2: ", day1part2.Run(readFile("day1/input/data")))

	fmt.Println("\nDay 2")
	fmt.Println("Part 1: ", day2part1.Run(readFile("day2/input/data")))
	fmt.Println("Part 2: ", day2part2.Run(readFile("day2/input/data")))

	fmt.Println("\nDay 3")
	fmt.Println("Part 1: ", day3part1.Run(readFile("day3/input/data")))
}
