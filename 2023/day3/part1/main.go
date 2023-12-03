package part1

import (
	"regexp"
	"strconv"
)

type Symbol struct {
	startIndex int
	endIndex   int
}

type PartNumber struct {
	number     int
	startIndex int
	endIndex   int
}

func Run(input []string) int {
	var sum int

	for i, line := range input {
		var nearSymbols []Symbol

		currentNumbers := findNumbers(line)
		nearSymbols = append(nearSymbols, findSpecialCharacters(line)...)

		if i != 0 {
			nearSymbols = append(nearSymbols, findSpecialCharacters(input[i-1])...)
		}

		if (len(input) - 1) > i {
			nearSymbols = append(nearSymbols, findSpecialCharacters(input[i+1])...)
		}

		for _, number := range currentNumbers {
			if number.isValid(nearSymbols) {
				sum = sum + number.number
			}
		}
	}

	return sum
}

func (partNumber PartNumber) isValid(nearSymbols []Symbol) bool {
	for _, symbol := range nearSymbols {
		if symbol.isAdjacentTo(partNumber) {
			return true
		}
	}

	return false
}

func (symbol Symbol) isAdjacentTo(partNumber PartNumber) bool {
	validStart := partNumber.startIndex - 1
	validEnd := partNumber.endIndex + 1

	if symbol.startIndex >= validStart && symbol.startIndex <= validEnd {
		return true
	}

	if symbol.endIndex >= validStart && symbol.endIndex <= validEnd {
		return true
	}

	return false
}

func findNumbers(line string) []PartNumber {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringIndex(line, -1)

	var partNumbers []PartNumber
	for _, match := range matches {
		startIndex, endIndex := getMatch(match)
		number, _ := strconv.Atoi(line[startIndex:endIndex])

		partNumbers = append(partNumbers, PartNumber{number, startIndex, endIndex - 1})
	}

	return partNumbers
}

func findSpecialCharacters(line string) []Symbol {
	re := regexp.MustCompile(`[^A-Za-z0-9\.]`)
	matches := re.FindAllStringIndex(line, -1)

	var symbols []Symbol
	for _, match := range matches {
		symbols = append(symbols, Symbol{match[0], match[1] - 1})
	}

	return symbols
}

func getMatch(match []int) (int, int) {
	return match[0], match[1]
}
