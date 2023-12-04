package part1

import (
	"testing"
)

func TestGetWinningNumberCount(t *testing.T) {
	t.Run("with 2 winning numbers", func(t *testing.T) {
		expected := 2
		input := ScratchCard{1, []int{1, 2, 3, 4, 5}, []int{4, 5}}

		result := input.GetWinningNumberCount()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("with 3 winning numbers", func(t *testing.T) {
		expected := 3
		input := ScratchCard{1, []int{1, 2, 3, 4, 5}, []int{3, 4, 5}}

		result := input.GetWinningNumberCount()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("with 5 winning numbers", func(t *testing.T) {
		expected := 5
		input := ScratchCard{1, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}

		result := input.GetWinningNumberCount()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
