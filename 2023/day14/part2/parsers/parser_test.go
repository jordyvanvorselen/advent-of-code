package parsers

import (
	"2023/day14/part2/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expected := []models.RockPositionList{
		{
			Rows: []string{
				"#.##O.##.",
				"O.#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#O",
			},
			Columns: []string{
				"#O##..#",
				"..##...",
				"##..###",
				"#....#.",
				"O#..#.#",
				".#..#.#",
				"#....#.",
				"##..###",
				"..##..O",
			},
		},
	}

	input := []string{
		"#.##O.##.",
		"O.#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#O",
	}

	result := Parse(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v but expected %+v", result, expected)
	}
}
