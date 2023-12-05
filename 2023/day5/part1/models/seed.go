package models

type Seed struct {
	Number int
}

func (s Seed) FindLocation(categoryMappings []CategoryMapping) int {
	sourceForNextMapping := s.Number

	for _, cm := range categoryMappings {
		sourceForNextMapping = cm.FindDestination(sourceForNextMapping)
	}

	return sourceForNextMapping
}
