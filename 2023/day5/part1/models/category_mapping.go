package models

type CategoryMapping struct {
	From   Category
	To     Category
	Ranges []RangeMapping
}
