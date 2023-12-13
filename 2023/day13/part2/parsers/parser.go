package parsers

import (
	"2023/day13/part2/models"
	"github.com/samber/lo"
	"strings"
)

func Parse(lines []string) []models.MirrorList {
	var result []models.MirrorList
	var rows []string

	for i, line := range lines {
		if i+1 >= len(lines) {
			rows = append(rows, line)
		}

		if line == "" || i+1 >= len(lines) {
			columns := getColumnsForRows(rows)
			result = append(result, models.MirrorList{Rows: rows, Columns: columns})
			rows = []string{}
			continue
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
