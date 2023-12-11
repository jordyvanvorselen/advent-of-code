package part2

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("oh wow, galaxies!", func(t *testing.T) {
		expected := 1030
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

		result := Run(input, 10)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("with factor 100", func(t *testing.T) {
		expected := 8410
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

		result := Run(input, 100)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
