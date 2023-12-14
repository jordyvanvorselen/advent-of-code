package utils

import (
	"github.com/samber/lo"
	"strings"
)

func RowsToColumns(columns []string) []string {
	return ColumnsToRows(columns)
}

func ColumnsToRows(rows []string) []string {
	lineLength := len(rows[0])
	var result []string

	for i := 0; i < lineLength; i++ {
		result = append(result, strings.Join(lo.Map(rows, func(row string, index int) string {
			return string(row[i])
		}), ""))
	}

	return result
}
