package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("2023/day-1/input/data")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	var calibrationValues []int
	foundFirst := false
	first := ""
	last := ""

	for _, line := range lines {
		for index, character := range line {
			if unicode.IsDigit(character) {
				if !foundFirst {
					first = string(character)
					last = string(character)

					foundFirst = true
				} else {
					last = string(character)
				}
			} else {
				digit := tryConvertDigit(line[index:])

				if digit == "" {
					continue
				}

				if !foundFirst {
					first = digit
					last = digit

					foundFirst = true
				} else {
					last = digit
				}
			}
		}

		finalValue, _ := strconv.Atoi(first + last)
		calibrationValues = append(calibrationValues, finalValue)

		foundFirst = false
	}

	sum := 0
	for _, num := range calibrationValues {
		sum += num
	}

	fmt.Println("Sum of calibration values: ", sum)
}

var spelledOutDigits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func tryConvertDigit(potentialDigit string) string {
	for i, actualDigit := range spelledOutDigits {
		if strings.Index(potentialDigit, actualDigit) == 0 {
			return strconv.Itoa(i + 1)
		}
	}

	return ""
}
