package parsers

import (
	"2023/day6/part2/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedRecord := models.Record{
		Time: 71530, Distance: 940200,
	}
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	resultRecord := Parse(input)

	if !reflect.DeepEqual(resultRecord, expectedRecord) {
		t.Errorf("got %+v but expected %+v", resultRecord, expectedRecord)
	}
}
