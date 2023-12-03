package part1

import (
	"regexp"
	"strconv"
)

func Run(lines []string) int {
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

	return sum
}
