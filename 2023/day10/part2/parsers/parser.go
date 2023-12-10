package parsers

import (
	"2023/day10/part2/models"
	"strings"
)

func Parse(lines []string) (int, models.Graph, int) {
	//result := make(map[int][]models.Edge)
	var result models.Graph
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
				result = append(result, models.Vertex{Index: index, Edges: []models.Edge{models.EdgeToLeft(index), models.EdgeToRight(index)}})
			case "7":
				result = append(result, models.Vertex{Index: index, Edges: []models.Edge{models.EdgeToLeft(index), models.EdgeToBottom(index, lineLength)}})
			case "|":
				result = append(result, models.Vertex{Index: index, Edges: []models.Edge{models.EdgeToTop(index, lineLength), models.EdgeToBottom(index, lineLength)}})
			case "J":
				result = append(result, models.Vertex{Index: index, Edges: []models.Edge{models.EdgeToTop(index, lineLength), models.EdgeToLeft(index)}})
			case "L":
				result = append(result, models.Vertex{Index: index, Edges: []models.Edge{models.EdgeToTop(index, lineLength), models.EdgeToRight(index)}})
			case "F":
				result = append(result, models.Vertex{Index: index, Edges: []models.Edge{models.EdgeToBottom(index, lineLength), models.EdgeToRight(index)}})
			case ".":
				result = append(result, models.Vertex{Index: index, Edges: []models.Edge{}})
			case "S":
				charBelow := lines[lineIndex+1][charIndex]
				var charAbove uint8
				var charLeft uint8
				var charRight uint8

				if charIndex > lineLength {
					charAbove = lines[lineIndex-1][charIndex]
				}

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

				result = append(result, models.Vertex{Index: index, Edges: edges})
				start = index
			}
		}
	}

	return start, result, lineLength
}
