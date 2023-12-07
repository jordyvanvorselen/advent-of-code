package models

import (
	"2023/day7/part1/utils"
	"unicode"
)

type Card struct {
	Label string
}

func (c Card) Strength() int {
	if char := rune(c.Label[0]); unicode.IsDigit(char) {
		return utils.ToInt(c.Label) - 2
	}

	others := []string{"T", "J", "Q", "K", "A"}

	for i, o := range others {
		if c.Label == o {
			return 8 + i
		}
	}

	panic("Not a valid card!")
}

func (c Card) isStrongerThan(other Card) int {
	if c.Strength() > other.Strength() {
		return 1
	}

	if c.Strength() == other.Strength() {
		return 0
	}

	return -1
}
