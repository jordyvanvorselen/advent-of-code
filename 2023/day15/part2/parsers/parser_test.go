package parsers

import (
	"2023/day15/part2/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expected := []models.InitItem{
		{"rn", models.ADD, 1},
		{"cm", models.REMOVE, 0},
		{"qp", models.ADD, 3},
		{"cm", models.ADD, 2},
		{"qp", models.REMOVE, 0},
		{"pc", models.ADD, 4},
		{"ot", models.ADD, 9},
		{"ab", models.ADD, 5},
		{"pc", models.REMOVE, 0},
		{"pc", models.ADD, 6},
		{"ot", models.ADD, 7},
	}

	input := []string{
		"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
	}

	result := Parse(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v but expected %+v", result, expected)
	}
}
