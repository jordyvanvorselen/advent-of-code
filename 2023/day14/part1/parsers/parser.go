package parsers

import (
	"2023/day14/part1/models"
	"github.com/samber/lo"
	"strings"
)

func Parse(lines []string) []models.RockPositionList {
	var result []models.RockPositionList
	var rows []string

	for i, line := range lines {
		if i+1 >= len(lines) {
			rows = append(rows, line)
		}

		if i+1 >= len(lines) {
			columns := getColumnsForRows(rows)
			result = append(result, models.RockPositionList{Rows: rows, Columns: columns})
			rows = []string{}
			break
		}

		rows = append(rows, line)
	}

	return result
}

func getColumnsForRows(rows []string) []string {
	lineLength := len(rows[0])
	var result []string

	for i := 0; i < lineLength; i++ {
		result = append(result, strings.Join(lo.Map(rows, func(row string, index int) string {
			return string(row[i])
		}), ""))
	}

	return result
}
