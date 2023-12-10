package part1

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("oasis' are nice, but did you ever see so many beautiful metal pipes?", func(t *testing.T) {
		expected := 4
		input := []string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}

	})

	t.Run("even more beautiful", func(t *testing.T) {
		expected := 4
		input := []string{
			"-L|F7",
			"7S-7|",
			"L|7||",
			"-L-J|",
			"L|-JF",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}

	})

	t.Run("the most beautiful", func(t *testing.T) {
		expected := 8
		input := []string{
			"7-F7-",
			".FJ|7",
			"SJLL7",
			"|F--J",
			"LJ.LJ",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}

	})
}
