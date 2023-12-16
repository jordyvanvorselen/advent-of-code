package part1

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("apparently I'm coding in a cave now, with some random mirrors and splitter thingies. cool", func(t *testing.T) {
		expected := 10
		input := []string{
			`\|...\....`,
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
