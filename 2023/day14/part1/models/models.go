package models

import (
	"2023/day14/part1/utils"
)

type RockPositionList struct {
	Rows    []string
	Columns []string
}

type Direction string

const (
	NORTH Direction = "N"
	WEST  Direction = "W"
	SOUTH Direction = "S"
	EAST  Direction = "E"
)

func (rpl RockPositionList) Cycle() RockPositionList {
	north := rpl.Tilt(NORTH)
	west := north.Tilt(WEST)
	south := west.Tilt(SOUTH)
	east := south.Tilt(EAST)

	return east
}

func (rpl RockPositionList) Tilt(direction Direction) RockPositionList {
	var lines []string
	var start []string

	if direction == NORTH || direction == SOUTH {
		start = rpl.Columns
	} else {
		start = rpl.Rows
	}

	for _, line := range start {
		var newLine string

		if direction == NORTH || direction == WEST {
			newLine = line
		} else {
			newLine = utils.String(line).Reverse()
		}

		for charIdx := 0; charIdx < len(newLine); charIdx++ {
			var nextChar string

			if charIdx == 0 {
				continue
			}

			char := string(newLine[charIdx])

			nextChar = string(newLine[charIdx-1])

			if char == "O" && nextChar == "." {
				shiftPlaces := freeFromIndex(newLine, charIdx)
				newLine = utils.String(newLine).ReplaceAt(charIdx, ".")
				newLine = utils.String(newLine).ReplaceAt(charIdx-shiftPlaces, "O")
			}
		}

		if direction == NORTH || direction == WEST {
			lines = append(lines, newLine)
		} else {
			lines = append(lines, utils.String(newLine).Reverse())
		}
	}

	if direction == NORTH || direction == SOUTH {
		return RockPositionList{Rows: utils.ColumnsToRows(lines), Columns: lines}
	} else {
		return RockPositionList{Rows: lines, Columns: utils.RowsToColumns(lines)}
	}

}

func freeFromIndex(line string, index int) int {
	var result int

	// column:  N  ...O...#...  S
	// row:     W  ...O...#...  E
	for _, c := range utils.String(line[:index]).Reverse() {
		if string(c) == "." {
			result++
			continue
		}

		break
	}

	return result
}
