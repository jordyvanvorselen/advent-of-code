package parsers

import (
	"2023/day7/part1/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedRecords := models.Hands{
		{
			Cards: []models.Card{
				{"3"}, {"2"}, {"T"}, {"3"}, {"K"},
			},
			Bid: 765,
		},
		{
			Cards: []models.Card{
				{"K"}, {"T"}, {"J"}, {"J"}, {"T"},
			},
			Bid: 220,
		},
	}
	input := []string{
		"32T3K 765",
		"KTJJT 220",
	}

	resultRecords := Parse(input)

	if !reflect.DeepEqual(resultRecords, expectedRecords) {
		t.Errorf("got %+v but expected %+v", resultRecords, expectedRecords)
	}
}
