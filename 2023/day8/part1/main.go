package part1

import (
	"2023/day8/part1/parsers"
	"2023/day8/part1/utils"
	"github.com/emirpasic/gods/trees/avltree"
	"strings"
)

func Run(input []string) int {
	var result int

	instructions, tree := parsers.Parse(input)

	instructions = strings.ReplaceAll(instructions, "L", "0")
	instructions = strings.ReplaceAll(instructions, "R", "1")

	traverseTree(instructions, tree, "AAA", &result)

	return result
}

func traverseTree(instructions string, tree *avltree.Tree, curr string, result *int) {
	if curr == "ZZZ" {
		return
	}

	instrIdx := *result % len(instructions)
	instruction := utils.ToInt(instructions[instrIdx : instrIdx+1])

	next, _ := tree.Get(curr)

	if nextVal, ok := next.([]string); ok {
		*result = *result + 1
		traverseTree(instructions, tree, nextVal[instruction], result)
	} else {
		panic("Invalid next value!")
	}
}
