package part2

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("is hitting the reindeer considered animal abuse or self defence in this case?", func(t *testing.T) {
		expected := 51
		input := []string{
			`.|...\....`,
			`|.-.\.....`,
			`.....|-...`,
			`........|.`,
			`..........`,
			`.........\`,
			`..../.\\..`,
			`.-.-/..|..`,
			`.|....-|.\`,
			`..//.|....`,
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
