package models

type CategoryMapping struct {
	From   Category
	To     Category
	Ranges []RangeMapping
}

func (cm CategoryMapping) FindDestination(source int) int {
	for _, rm := range cm.Ranges {
		destination := rm.FindDestination(source)

		if destination != source {
			return destination
		}
	}

	return source
}
