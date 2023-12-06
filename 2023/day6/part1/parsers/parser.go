package parsers

import (
	"2023/day6/part1/models"
	"strconv"
	"strings"
)

func Parse(lines []string) []models.Record {
	var result []models.Record
	var times []int
	var distances []int

	maybeTimes := strings.Fields(lines[0])
	maybeDistances := strings.Fields(lines[1])

	for _, word := range maybeTimes {
		if number, err := strconv.Atoi(word); err == nil {
			times = append(times, number)
		}
	}

	for _, word := range maybeDistances {
		if number, err := strconv.Atoi(word); err == nil {
			distances = append(distances, number)
		}
	}

	for i, _ := range times {
		result = append(result, models.Record{Time: times[i], Distance: distances[i]})
	}

	return result
}
