package part2

import (
	"2023/day8/part2/parsers"
	"2023/day8/part2/utils"
	"github.com/emirpasic/gods/trees/avltree"
	"github.com/samber/lo"
	"strings"
)

func Run(input []string) int {

	instructions, tree := parsers.Parse(input)

	instructions = strings.ReplaceAll(instructions, "L", "0")
	instructions = strings.ReplaceAll(instructions, "R", "1")

	currs := lo.FilterMap(tree.Keys(), func(key interface{}, index int) (string, bool) {
		if val, ok := key.(string); ok {
			if strings.HasSuffix(val, "A") {
				return val, true
			}

			return "", false
		} else {
			panic("Invalid key!")
		}
	})

	var steps []int
	for _, curr := range currs {
		var result int
		var stop bool

		for {
			curr, stop = traverseTree(instructions, tree, curr, &result)

			if stop {
				break
			}
		}

		steps = append(steps, result)
	}

	return lcm(steps[0], steps[1], steps[2:]...)
}

func traverseTree(instructions string, tree *avltree.Tree, curr string, result *int) (string, bool) {
	if strings.HasSuffix(curr, "Z") {
		return "", true
	}

	instrIdx := *result % len(instructions)
	instruction := utils.ToInt(instructions[instrIdx : instrIdx+1])

	next, _ := tree.Get(curr)

	if nextVal, ok := next.([]string); ok {
		*result = *result + 1
		return nextVal[instruction], false
	}

	panic("Invalid next value!")
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
