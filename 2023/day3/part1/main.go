package part1

import (
	"regexp"
	"strconv"
)

type Indexable struct {
	startIndex int
	endIndex   int
}

type Symbol struct {
	Indexable
}

type PartNumber struct {
	number int
	Indexable
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

func (p PartNumber) isValid(nearSymbols []Symbol) bool {
	for _, symbol := range nearSymbols {
		if symbol.isAdjacentTo(p.Indexable) {
			return true
		}
	}

	return false
}

func (i Indexable) isAdjacentTo(i2 Indexable) bool {
	validStart := i2.startIndex - 1
	validEnd := i2.endIndex + 1

	if i.startIndex >= validStart && i.startIndex <= validEnd {
		return true
	}

	if i.endIndex >= validStart && i.endIndex <= validEnd {
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

		partNumbers = append(partNumbers, PartNumber{number, Indexable{startIndex, endIndex - 1}})
	}

	return partNumbers
}

func findSpecialCharacters(line string) []Symbol {
	re := regexp.MustCompile(`[^A-Za-z0-9\.]`)
	matches := re.FindAllStringIndex(line, -1)

	var symbols []Symbol
	for _, match := range matches {
		symbols = append(symbols, Symbol{Indexable{match[0], match[1] - 1}})
	}

	return symbols
}

func getMatch(match []int) (int, int) {
	return match[0], match[1]
}
