package main

import (
	day1part1 "2023/day1/part1"
	day1part2 "2023/day1/part1"
	day10part1 "2023/day10/part1"
	day10part2 "2023/day10/part2"
	day11part1 "2023/day11/part1"
	day11part2 "2023/day11/part2"
	day12part1 "2023/day12/part1"
	day2part1 "2023/day2/part1"
	day2part2 "2023/day2/part1"
	day3part1 "2023/day3/part1"
	day3part2 "2023/day3/part1"
	day4part1 "2023/day4/part1"
	day4part2 "2023/day4/part1"
	day5part1 "2023/day5/part1"
	day5part2 "2023/day5/part1"
	day6part1 "2023/day6/part1"
	day6part2 "2023/day6/part1"
	day7part1 "2023/day7/part1"
	day7part2 "2023/day7/part1"
	day8part1 "2023/day8/part1"
	day8part2 "2023/day8/part1"
	day9part1 "2023/day9/part1"
	day9part2 "2023/day9/part1"
	//day12part2 "2023/day12/part2"
	day13part1 "2023/day13/part1"
	day13part2 "2023/day13/part2"
	day14part1 "2023/day14/part1"
	day14part2 "2023/day14/part2"
	day15part1 "2023/day15/part1"
	day15part2 "2023/day15/part2"
	day16part1 "2023/day16/part1"
	day16part2 "2023/day16/part2"
	"fmt"
)

func main() {
	if false {
		fmt.Println("Day 1")
		fmt.Println("Part 1: ", day1part1.Run(readFile("day1/input/data")))
		fmt.Println("Part 2: ", day1part2.Run(readFile("day1/input/data")))

		fmt.Println("\nDay 2")
		fmt.Println("Part 1: ", day2part1.Run(readFile("day2/input/data")))
		fmt.Println("Part 2: ", day2part2.Run(readFile("day2/input/data")))

		fmt.Println("\nDay 3")
		fmt.Println("Part 1: ", day3part1.Run(readFile("day3/input/data")))
		fmt.Println("Part 2: ", day3part2.Run(readFile("day3/input/data")))

		fmt.Println("\nDay 4")
		fmt.Println("Part 1: ", day4part1.Run(readFile("day4/input/data")))
		fmt.Println("Part 2: ", day4part2.Run(readFile("day4/input/data")))

		fmt.Println("\nDay 5")
		fmt.Println("Part 1: ", day5part1.Run(readFile("day5/input/data")))
		fmt.Println("Part 2: ", day5part2.Run(readFile("day5/input/data")))

		fmt.Println("\nDay 6")
		fmt.Println("Part 1: ", day6part1.Run(readFile("day6/input/data")))
		fmt.Println("Part 2: ", day6part2.Run(readFile("day6/input/data")))

		fmt.Println("\nDay 7")
		fmt.Println("Part 1: ", day7part1.Run(readFile("day7/input/data")))
		fmt.Println("Part 2: ", day7part2.Run(readFile("day7/input/data")))

		fmt.Println("\nDay 8")
		fmt.Println("Part 1: ", day8part1.Run(readFile("day8/input/data")))
		fmt.Println("Part 2: ", day8part2.Run(readFile("day8/input/data")))

		fmt.Println("\nDay 9")
		fmt.Println("Part 1: ", day9part1.Run(readFile("day9/input/data")))
		fmt.Println("Part 2: ", day9part2.Run(readFile("day9/input/data")))

		fmt.Println("\nDay 10")
		fmt.Println("Part 1: ", day10part1.Run(readFile("day10/input/data")))
		fmt.Println("Part 2: ", day10part2.Run(readFile("day10/input/data")))

		fmt.Println("\nDay 11")
		fmt.Println("Part 1: ", day11part1.Run(readFile("day11/input/data")))
		fmt.Println("Part 2: ", day11part2.Run(readFile("day11/input/data"), 1000000))

		fmt.Println("\nDay 12")
		fmt.Println("Part 1: ", day12part1.Run(readFile("day12/input/data")))
		//fmt.Println("Part 2: ", day12part2.Run(readFile("day12/input/data")))

		fmt.Println("\nDay 13")
		fmt.Println("Part 1: ", day13part1.Run(readFile("day13/input/data")))
		fmt.Println("Part 2: ", day13part2.Run(readFile("day13/input/data")))

		fmt.Println("\nDay 14")
		fmt.Println("Part 1: ", day14part1.Run(readFile("day14/input/data")))
		fmt.Println("Part 2: ", day14part2.Run(readFile("day14/input/data")))

		fmt.Println("\nDay 15")
		fmt.Println("Part 1: ", day15part1.Run(readFile("day15/input/data")))
		fmt.Println("Part 2: ", day15part2.Run(readFile("day15/input/data")))
	}

	fmt.Println("\nDay 16")
	fmt.Println("Part 1: ", day16part1.Run(readFile("day16/input/data")))
	fmt.Println("Part 2: ", day16part2.Run(readFile("day16/input/data")))
}
