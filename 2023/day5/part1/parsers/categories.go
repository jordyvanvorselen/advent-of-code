package parsers

import (
	"2023/day5/part1/models"
	"regexp"
)

func ParseCategories(line string) (models.Category, models.Category) {
	pattern := regexp.MustCompile(`(\w+)-to-(\w+) map:`)
	matches := pattern.FindStringSubmatch(line)

	return models.Category(matches[1]), models.Category(matches[2])
}
