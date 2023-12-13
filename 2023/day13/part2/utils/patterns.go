package utils

import "regexp"

func findAllByPattern(line string, pattern string) []string {
	regex := regexp.MustCompile(pattern)

	return regex.FindAllString(line, -1)
}

func findFirstByPattern(line string, fromIndex int, pattern string) string {
	regex := regexp.MustCompile(pattern)

	return regex.FindString(line[fromIndex:])
}
