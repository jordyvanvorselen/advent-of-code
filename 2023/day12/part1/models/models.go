package models

import (
	"slices"
)

type ConditionRecord struct {
	Raw    []string
	Broken []int
}

func (c ConditionRecord) PossibleBrokenCombinations() int {
	var result int
	doesMatch := false
	startPositions := []int{0}

	for i, _ := range c.Broken {
		for _, start := range startPositions {
			doesMatch, startPositions = c.brokenListDoesMatch(c.Broken[i:], start)

			if doesMatch {
				result += len(startPositions)
			}
		}
	}

	return result
}

func (c ConditionRecord) brokenListDoesMatch(broken []int, startPosition int) (bool, []int) {
	var newStartPositions []int

	if startPosition >= len(c.Raw) {
		return false, []int{}
	}

	raw := c.Raw[startPosition:]

	for brokenIdx, b := range broken {
		brokenDidMatch := false

		for charIdx := 0; charIdx < len(raw)-b+1; charIdx++ {
			var charLeft string
			var charRight string

			chars := c.Raw[charIdx : charIdx+b]

			if charIdx != 0 {
				charLeft = c.Raw[charIdx-1]
			}

			if charIdx < len(c.Raw)-b {
				charRight = c.Raw[charIdx+b]
			}

			if fitsInSpot(charLeft, charRight, chars) {
				if brokenIdx == 0 {
					newStartPositions = append(newStartPositions, startPosition+charIdx+b+1)
				}

				brokenDidMatch = true
			}
		}

		if !brokenDidMatch {
			return false, []int{}
		}
	}

	return true, newStartPositions
}

func fitsInSpot(charLeft string, charRight string, chars []string) bool {
	return charLeft != "#" && charRight != "#" && !slices.Contains(chars, ".")
}
