package models

type Seed struct {
	Number int
}

func (s Seed) FindLocation(rangeMappings []RangeMapping) int {
	var location int

	for _, rm := range rangeMappings {
		destination := rm.FindDestination(s)
	}

	return location
}
