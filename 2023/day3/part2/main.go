package part2

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
		var nearPartNumbers []PartNumber

		currentSymbols := findSpecialCharacters(line)
		nearPartNumbers = append(nearPartNumbers, findNumbers(line)...)

		if i != 0 {
			nearPartNumbers = append(nearPartNumbers, findNumbers(input[i-1])...)
		}

		if (len(input) - 1) > i {
			nearPartNumbers = append(nearPartNumbers, findNumbers(input[i+1])...)
		}

		for _, symbol := range currentSymbols {
			sum = sum + symbol.getGearRatio(nearPartNumbers)
		}
	}

	return sum
}

func (s Symbol) getGearRatio(nearPartNumbers []PartNumber) int {
	var validNumbers []int

	for _, number := range nearPartNumbers {
		if number.isAdjacentTo(s.Indexable) {
			validNumbers = append(validNumbers, number.number)
		}
	}

	if len(validNumbers) == 2 {
		return validNumbers[0] * validNumbers[1]
	}

	return 0
}

func (i Indexable) isAdjacentTo(idx2 Indexable) bool {
	validStart := idx2.startIndex - 1
	validEnd := idx2.endIndex + 1

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
