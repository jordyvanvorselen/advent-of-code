package utils

import "regexp"

func RemoveNonDigits(input string) string {
	reg := regexp.MustCompile("[^0-9]")

	return reg.ReplaceAllString(input, "")
}
