package parsers

import (
	"2023/day12/part1/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expected := []models.ConditionRecord{
		{
			Raw:    []string{"#", ".", "#", ".", "#", "#", "#"},
			Broken: []models.Group{1, 1, 3},
		},
	}

	input := []string{
		"#.#.### 1,1,3",
	}

	result := Parse(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v but expected %+v", result, expected)
	}
}
