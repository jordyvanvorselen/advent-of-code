package part1

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("oh wow, galaxies!", func(t *testing.T) {
		expected := 374
		input := []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
