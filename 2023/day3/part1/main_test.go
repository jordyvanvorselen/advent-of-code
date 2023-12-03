package part1

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 4361
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	result := Run(input)

	if result != expected {
		t.Errorf("got %d but expected %d", result, expected)
	}
}
