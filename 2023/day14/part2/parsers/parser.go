package parsers

import (
	"2023/day14/part2/models"
	"2023/day14/part2/utils"
)

func Parse(lines []string) models.RockPositionList {
	var result models.RockPositionList
	var rows []string

	for i, line := range lines {
		if i+1 >= len(lines) {
			rows = append(rows, line)
		}

		if i+1 >= len(lines) {
			columns := utils.ColumnsToRows(rows)
			result = models.RockPositionList{Rows: rows, Columns: columns}
			rows = []string{}
			break
		}

		rows = append(rows, line)
	}

	return result
}
