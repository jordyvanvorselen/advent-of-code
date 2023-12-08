package part2

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("simultaneous traversing through instructions", func(t *testing.T) {
		expected := 6
		input := []string{
			"LR",
			"",
			"11A = (11B, XXX)",
			"11B = (XXX, 11Z)",
			"11Z = (11B, XXX)",
			"22A = (22B, XXX)",
			"22B = (22C, 22C)",
			"22C = (22Z, 22Z)",
			"22Z = (22B, 22B)",
			"XXX = (XXX, XXX)",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}

	})
}
