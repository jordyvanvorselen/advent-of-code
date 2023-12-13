package models

import (
	"2023/day13/part2/utils"
	"github.com/samber/lo"
)

type MirrorList struct {
	Rows    []string
	Columns []string
}

func (m MirrorList) FindPerfectRowReflectionIndex() (int, bool, bool) {
	orgIdx, orgMatch := findPerfectReflectionIndex(m.Rows, -1)

	for _, rows := range getAllCombinations(m.Rows) {
		idx, match := findPerfectReflectionIndex(rows, orgIdx)

		if match {
			return idx, true, false
		}
	}

	if orgMatch {
		return orgIdx, true, true
	}

	return -1, false, false
}

func (m MirrorList) FindPerfectColumnReflectionIndex() (int, bool, bool) {
	orgIdx, orgMatch := findPerfectReflectionIndex(m.Columns, -1)

	allPos := getAllCombinations(m.Columns)

	for _, columns := range allPos {
		idx, match := findPerfectReflectionIndex(columns, orgIdx)

		if match {
			return idx, true, false
		}
	}

	if orgMatch {
		return orgIdx, true, true
	}

	return -1, false, false
}

func getAllCombinations(lines []string) [][]string {
	var allCombinations [][]string

	for rowIdx, line := range lines {
		for charIdx, c := range line {
			var newChar string

			if string(c) == "#" {
				newChar = "."
			}

			if string(c) == "." {
				newChar = "#"
			}

			changedLine := utils.String(line).ReplaceAt(charIdx, newChar)
			newElement := lo.Replace(lines, lines[rowIdx], changedLine, -1)
			allCombinations = append(allCombinations, newElement)
		}
	}

	return allCombinations
}

func findPerfectReflectionIndex(lines []string, orgIdx int) (int, bool) {
	for i, line := range lines {
		if i+1 == len(lines) {
			break
		}

		nextLine := lines[i+1]

		if line == nextLine && i != orgIdx {
			if isPerfectReflection(lines, i) {
				return i, true
			}
		}
	}

	return -1, false
}

func isPerfectReflection(lines []string, index int) bool {
	indexA := index
	indexB := index + 1

	for {
		if indexA < 0 {
			break
		}

		if indexB >= len(lines) {
			break
		}

		if lines[indexA] != lines[indexB] {
			return false
		}

		indexA--
		indexB++
	}

	return true
}
