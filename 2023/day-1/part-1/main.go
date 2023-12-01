package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	for _, line := range lines {
		pattern := `^.*?(\d).*?(\d)[^\d]*$|(\d)`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(line)

		if matches[1] != "" {
			calibrationValue, _ := strconv.Atoi(matches[1] + matches[2])
			calibrationValues = append(calibrationValues, calibrationValue)
		} else {
			calibrationValue, _ := strconv.Atoi(matches[3] + matches[3])
			calibrationValues = append(calibrationValues, calibrationValue)
		}
	}

	sum := 0
	for _, num := range calibrationValues {
		sum += num
	}

	fmt.Println("Sum of calibration values: ", sum)
}
