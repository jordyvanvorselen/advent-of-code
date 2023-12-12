package models

import (
	"slices"
)

type ConditionRecord struct {
	Raw    []string
	Broken []Group
}

type Group int

func (c ConditionRecord) PossibleBrokenCombinations() int {
	var startIndex int
	var result int

	for {
		fits := false
		fits, startIndex = c.CombinationIsPossible(startIndex)

		if !fits {
			break
		}

		result++
	}

	return result
}

func (c ConditionRecord) CombinationIsPossible(startIdx int) (bool, int) {
	newStartIdx := -1

	startIdxForNextChar := startIdx

	for brokenIdx, b := range c.Broken {
		rawLeft := c.Raw[startIdxForNextChar:]
		fitsInSomewhere := false

		for charIdx, _ := range rawLeft {
			if b.fitsInGap(rawLeft, charIdx) {
				startIdxForNextChar = startIdxForNextChar + charIdx + int(b) + 1
				fitsInSomewhere = true

				if newStartIdx == -1 {
					newStartIdx = startIdx + charIdx + int(b)
				}

				break
			}
		}

		if !fitsInSomewhere {
			break
		}

		if brokenIdx == len(c.Broken)-1 {
			return true, newStartIdx
		}
	}

	return false, 0
}

func (g Group) fitsInGap(raw []string, index int) bool {
	if len(raw) < index+int(g) {
		return false
	}

	gap := raw[index : index+int(g)]

	if slices.Contains(gap, ".") {
		return false
	}

	// if there is a character after the gap
	if len(raw) > index+int(g)+1 {
		maybeDefect := raw[index+int(g) : index+int(g)+1]
		if maybeDefect[0] == "#" {
			return false
		}
	}

	return true
}
