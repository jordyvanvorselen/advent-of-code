package parsers

import (
	"2023/day5/part1/models"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedSeeds := []models.Seed{{79}, {14}, {55}, {13}}
	expectedCategoryMappings := []models.CategoryMapping{
		{
			From: models.Category("seed"),
			To:   models.Category("soil"),
			Ranges: []models.RangeMapping{
				{DestinationRangeStart: 50, SourceRangeStart: 98, RangeLength: 2},
				{DestinationRangeStart: 52, SourceRangeStart: 50, RangeLength: 48},
			},
		},
	}
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	}

	resultSeeds, resultCategoryMappings := Parse(input)

	if !reflect.DeepEqual(resultSeeds, expectedSeeds) {
		t.Errorf("got %+v but expected %+v", resultSeeds, expectedSeeds)
	}

	if !reflect.DeepEqual(resultCategoryMappings, expectedCategoryMappings) {
		t.Errorf("got %+v but expected %+v", resultCategoryMappings, expectedCategoryMappings)
	}
}
