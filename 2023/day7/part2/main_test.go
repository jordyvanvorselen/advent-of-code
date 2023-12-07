package part2

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 5905
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	result := Run(input)

	if result != expected {
		t.Errorf("got %d but expected %d", result, expected)
	}
}
