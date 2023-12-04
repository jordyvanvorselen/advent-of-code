package part1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parse(lines []string) []ScratchCard {
	var scratchCards []ScratchCard

	for _, line := range lines {
		inputString := line

		pattern := fmt.Sprintf("(?i)%s", `Card \d\: `) // (?i) for case-insensitive matching
		re := regexp.MustCompile(pattern)
		result := re.ReplaceAllString(inputString, "")

		splitCard := strings.Split(result, "|")

		scratchCards = append(scratchCards, ScratchCard{winning: getNumbers(splitCard[0]), numbers: getNumbers(splitCard[1])})
	}

	return scratchCards
}

func getNumbers(input string) []int {
	var result []int
	numbers := strings.Fields(input)

	for _, n := range numbers {
		i, _ := strconv.Atoi(n)
		result = append(result, i)
	}

	return result
}
