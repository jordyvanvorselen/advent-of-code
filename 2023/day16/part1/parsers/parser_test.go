package parsers

import (
	"2023/day16/part1/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expected := models.Grid{
		"0,0": models.Tile{X: 0, Y: 0, Type: models.EMPTY},
		"1,0": models.Tile{X: 1, Y: 0, Type: models.VERTICAL_SPLITTER},
		"2,0": models.Tile{X: 2, Y: 0, Type: models.EMPTY},
		"0,1": models.Tile{X: 0, Y: 1, Type: models.VERTICAL_SPLITTER},
		"1,1": models.Tile{X: 1, Y: 1, Type: models.EMPTY},
		"2,1": models.Tile{X: 2, Y: 1, Type: models.HORIZONTAL_SPLITTER},
	}

	input := []string{
		`.|.`,
		`|.-`,
	}

	result := Parse(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v but expected %+v", result, expected)
	}
}
