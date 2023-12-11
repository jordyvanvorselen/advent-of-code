package parsers

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/samber/lo"
	"slices"
	"strings"
)

func Parse(lines []string) (*dijkstra.Graph, []int) {
	result := dijkstra.NewGraph()
	var galaxies []int

	actualInput := expandGravitational(lines, len(lines[0]), len(lines))

	for lineIdx, line := range actualInput {
		fmt.Println(line)
		for charIdx, c := range line {
			index := lineIdx*len(line) + charIdx
			char := string(c)

			result.AddVertex(index)

			if char == "#" {
				galaxies = append(galaxies, index)
			}

			isTop := lineIdx == 0
			isBottom := lineIdx == len(actualInput)-1
			isLeft := charIdx == 0
			isRight := charIdx == len(line)-1

			if !isTop {
				result.AddArc(index, index-len(line), 1)
				result.AddArc(index-len(line), index, 1)
			}

			if !isBottom {
				result.AddArc(index, index+len(line), 1)
				result.AddArc(index+len(line), index, 1)
			}

			if !isLeft {
				result.AddArc(index, index-1, 1)
				result.AddArc(index-1, index, 1)
			}

			if !isRight {
				result.AddArc(index, index+1, 1)
				result.AddArc(index+1, index, 1)
			}
		}
	}

	return result, galaxies
}

func expandGravitational(lines []string, lineLength int, columnLength int) []string {
	result := lines
	var rowsToExpand []int
	var colsToExpand []int

	for i, line := range result {
		if strings.ReplaceAll(line, ".", "") == "" {
			rowsToExpand = append(rowsToExpand, i)
		}
	}

	for i := 0; i < columnLength; i++ {
		col := lo.Map(result, func(line string, index int) string {
			return string(line[i])
		})

		column := strings.Join(col, "")

		if strings.ReplaceAll(column, ".", "") == "" {
			colsToExpand = append(colsToExpand, i)
		}
	}

	rowsAdded := 0
	for _, i := range rowsToExpand {
		result = slices.Insert(result, i+rowsAdded, emptyLine(lineLength))
		// after adding a row, all others also need to take that new row into account
		rowsAdded++
	}

	colsAdded := 0
	for _, i := range colsToExpand {
		for lineIdx, _ := range result {
			line := result[lineIdx]
			slices.Replace(result, lineIdx, lineIdx+1, insertString(line, ".", i+colsAdded))
		}
		colsAdded++
	}

	return result
}

func emptyLine(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += "."
	}

	return result
}

func insertString(original, insert string, index int) string {
	return original[:index] + insert + original[index:]
}
