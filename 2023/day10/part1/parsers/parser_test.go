package parsers

import (
	"2023/day10/part1/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expected := map[int][]models.Edge{
		6:  {{11, 1}, {7, 1}},
		7:  {{6, 1}, {8, 1}},
		8:  {{7, 1}, {13, 1}},
		11: {{6, 1}, {16, 1}},
		13: {{8, 1}, {18, 1}},
		16: {{11, 1}, {17, 1}},
		17: {{16, 1}, {18, 1}},
		18: {{13, 1}, {17, 1}},
	}
	expectedStart := 6

	input := []string{
		".....", // 0 1 2 3 4
		".S-7.", // 5 6 7 8 9
		".|.|.", // 10 11 12 13 14
		".L-J.", // 15 16 17 18 19
		".....", // 20 21 22 23 24
	}

	start, result := Parse(input)

	if start != expectedStart {
		t.Errorf("got %d but expected %d", start, expectedStart)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v but expected %+v", result, expected)
	}
}
