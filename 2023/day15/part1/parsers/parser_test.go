package parsers

import (
	"2023/day15/part1/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expected := []models.InitItem{
		{"rn=1"},
		{"cm-"},
		{"qp=3"},
		{"cm=2"},
		{"qp-"},
		{"pc=4"},
		{"ot=9"},
		{"ab=5"},
		{"pc-"},
		{"pc=6"},
		{"ot=7"},
	}

	input := []string{
		"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
	}

	result := Parse(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v but expected %+v", result, expected)
	}
}
