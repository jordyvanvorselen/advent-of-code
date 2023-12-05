package models

import "slices"

type RangeMapping struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func (rm RangeMapping) FindDestination(seed Seed) int {
	if !slices.Contains(rm.sourceRange(), seed.Number) {
		return seed.Number
	}

	// WIP
	return 0
}

func (rm RangeMapping) destinationRange() []int {
	return createRange(rm.DestinationRangeStart, rm.DestinationRangeStart+rm.RangeLength)
}

func (rm RangeMapping) sourceRange() []int {
	return createRange(rm.SourceRangeStart, rm.SourceRangeStart+rm.RangeLength)
}

func createRange(start int, end int) []int {
	var result []int

	for i := start; i <= end; i++ {
		result = append(result, i)
	}

	return result
}
