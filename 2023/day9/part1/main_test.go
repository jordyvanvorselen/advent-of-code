package part1

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("without looping through instructions", func(t *testing.T) {
		expected := 114
		input := []string{
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}

	})
}
