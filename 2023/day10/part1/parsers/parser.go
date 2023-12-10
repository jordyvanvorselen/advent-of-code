package parsers

import (
	"2023/day10/part1/models"
	"strings"
)

func Parse(lines []string) (int, map[int][]models.Edge) {
	result := make(map[int][]models.Edge)
	lineLength := len(lines[0])
	var start int

	// "....."
	// ".S-7."
	// ".|.|."
	// ".L-J."
	// "....."
	for lineIndex, line := range lines {
		for charIndex, char := range line {
			index := lineIndex*lineLength + charIndex

			switch character := string(char); character {
			case "-":
				result[index] = []models.Edge{models.EdgeToLeft(index), models.EdgeToRight(index)}
			case "7":
				result[index] = []models.Edge{models.EdgeToLeft(index), models.EdgeToBottom(index, lineLength)}
			case "|":
				result[index] = []models.Edge{models.EdgeToTop(index, lineLength), models.EdgeToBottom(index, lineLength)}
			case "J":
				result[index] = []models.Edge{models.EdgeToTop(index, lineLength), models.EdgeToLeft(index)}
			case "L":
				result[index] = []models.Edge{models.EdgeToTop(index, lineLength), models.EdgeToRight(index)}
			case "F":
				result[index] = []models.Edge{models.EdgeToBottom(index, lineLength), models.EdgeToRight(index)}
			case "S":
				charAbove := lines[lineIndex-1][charIndex]
				charBelow := lines[lineIndex+1][charIndex]
				var charLeft uint8
				var charRight uint8

				if charIndex > 0 {
					charLeft = line[charIndex-1]
				}

				if charIndex < lineLength {
					charRight = line[charIndex+1]
				}

				var edges []models.Edge

				if strings.Contains("7|F", string(charAbove)) {
					edges = append(edges, models.EdgeToTop(index, lineLength))
				}

				if strings.Contains("|JL", string(charBelow)) {
					edges = append(edges, models.EdgeToBottom(index, lineLength))
				}

				if strings.Contains("-7J", string(charRight)) {
					edges = append(edges, models.EdgeToRight(index))
				}

				if strings.Contains("FL-", string(charLeft)) {
					edges = append(edges, models.EdgeToLeft(index))
				}

				result[index] = edges
				start = index
			}
		}
	}

	return start, result
}
