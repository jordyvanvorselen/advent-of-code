package part2

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("hey look even more mirrors! and some rocks...", func(t *testing.T) {
		expected := 64
		input := []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#....",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
