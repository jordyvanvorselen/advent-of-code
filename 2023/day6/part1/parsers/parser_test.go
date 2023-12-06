package parsers

import (
	"2023/day6/part1/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedRecords := []models.Record{
		{Time: 7, Distance: 9},
		{Time: 15, Distance: 40},
		{Time: 30, Distance: 200},
	}
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	resultRecords := Parse(input)

	if !reflect.DeepEqual(resultRecords, expectedRecords) {
		t.Errorf("got %+v but expected %+v", resultRecords, expectedRecords)
	}
}
