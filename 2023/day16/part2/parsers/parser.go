package parsers

import (
	"2023/day16/part2/models"
	"fmt"
)

func Parse(lines []string) models.Grid {
	result := make(models.Grid)

	for y, line := range lines {
		for x, c := range line {
			result[fmt.Sprintf("%d,%d", x, y)] = models.Tile{X: x, Y: y, Type: models.TileType(c)}
		}
	}

	return result
}
