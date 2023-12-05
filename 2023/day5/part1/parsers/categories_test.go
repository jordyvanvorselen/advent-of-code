package parsers

import (
	"2023/day5/part1/models"
	"testing"
)

func TestParseCategories(t *testing.T) {
	expectedFrom, expectedTo := models.Category("soil"), models.Category("water")
	input := "soil-to-water map:"

	from, to := ParseCategories(input)

	if from != expectedFrom {
		t.Errorf("got %s but expected %s", from, expectedFrom)
	}

	if to != expectedTo {
		t.Errorf("got %s but expected %s", to, expectedTo)
	}
}
