package parsers

import (
	"github.com/RyanCarrier/dijkstra"
	"github.com/samber/lo"
	"slices"
	"strings"
)

func Parse(lines []string, factor int) (*dijkstra.Graph, []int) {
	result := dijkstra.NewGraph()
	var galaxies []int

	rowsToIncrease, colsToIncrease := expandGravitational(lines, len(lines))

	for lineIdx, line := range lines {
		for charIdx, c := range line {
			index := lineIdx*len(line) + charIdx
			char := string(c)

			result.AddVertex(index)

			if char == "#" {
				galaxies = append(galaxies, index)
			}

			isTop := lineIdx == 0
			isBottom := lineIdx == len(lines)-1
			isLeft := charIdx == 0
			isRight := charIdx == len(line)-1

			rowDistance := 1
			colDistance := 1

			if slices.Contains(rowsToIncrease, lineIdx) {
				rowDistance = 1 * factor
			}

			if slices.Contains(colsToIncrease, charIdx) {
				colDistance = 1 * factor
			}

			if !isTop {
				result.AddArc(index, index-len(line), int64(rowDistance))
				result.AddArc(index-len(line), index, int64(rowDistance))
			}

			if !isBottom {
				result.AddArc(index, index+len(line), int64(rowDistance))
				result.AddArc(index+len(line), index, int64(rowDistance))
			}

			if !isLeft {
				result.AddArc(index, index-1, int64(colDistance))
				result.AddArc(index-1, index, int64(colDistance))
			}

			if !isRight {
				result.AddArc(index, index+1, int64(colDistance))
				result.AddArc(index+1, index, int64(colDistance))
			}
		}
	}

	return result, galaxies
}

func expandGravitational(lines []string, columns int) ([]int, []int) {
	result := lines
	var rowsToExpand []int
	var colsToExpand []int

	for i, line := range result {
		if strings.ReplaceAll(line, ".", "") == "" {
			rowsToExpand = append(rowsToExpand, i)
		}
	}

	for i := 0; i < columns; i++ {
		col := lo.Map(result, func(line string, index int) string {
			return string(line[i])
		})

		column := strings.Join(col, "")

		if strings.ReplaceAll(column, ".", "") == "" {
			colsToExpand = append(colsToExpand, i)
		}
	}

	return rowsToExpand, colsToExpand
}
