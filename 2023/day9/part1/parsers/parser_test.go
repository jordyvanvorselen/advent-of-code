package parsers

import (
	"2023/day9/part1/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expected := []models.Sequence{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}

	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	result := Parse(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v but expected %+v", result, expected)
	}
}
