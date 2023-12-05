package models

type RangeMapping struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func (rm RangeMapping) FindDestination(source int) int {
	if source < rm.SourceRangeStart || source >= rm.SourceRangeStart+rm.RangeLength {
		return source
	}

	return rm.DestinationRangeStart + (source - rm.SourceRangeStart)
}
