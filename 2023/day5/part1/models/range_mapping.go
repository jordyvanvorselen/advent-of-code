package models

import "slices"

type RangeMapping struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func (rm RangeMapping) FindDestination(source int) int {
	if !slices.Contains(rm.sourceRange(), source) {
		return source
	}

	return rm.destinationRange()[findIndex(source, rm.sourceRange())]
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

func findIndex(value int, list []int) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}

	return -1
}
