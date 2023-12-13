package part2

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("hey look some mirrors! today will probably be a fun day :')", func(t *testing.T) {
		expected := 400
		input := []string{
			"#.##..##.",
			"..#.##.#.",
			"##......#",
			"##......#",
			"..#.##.#.",
			"..##..##.",
			"#.#.##.#.",
			"",
			"#...##..#",
			"#....#..#",
			"..##..###",
			"#####.##.",
			"#####.##.",
			"..##..###",
			"#....#..#",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
