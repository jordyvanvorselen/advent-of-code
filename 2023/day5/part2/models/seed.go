package models

type Seeds struct {
	RangeStart  int
	RangeLength int
}

func (s Seeds) FindLocations(categoryMappings []CategoryMapping) []int {
	var result []int

	for i := 0; i < s.RangeLength; i++ {
		sourceForNextMapping := s.RangeStart + i

		for _, cm := range categoryMappings {
			sourceForNextMapping = cm.FindDestination(sourceForNextMapping)
		}

		result = append(result, sourceForNextMapping)
	}

	return result
}
