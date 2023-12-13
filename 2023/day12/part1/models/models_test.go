package models

import "testing"

func TestConditionRecord(t *testing.T) {
	t.Run("with gaps", func(t *testing.T) {
		expected := 4
		input := ConditionRecord{
			Raw:    []string{".", "?", "?", ".", ".", "?", "?", ".", ".", ".", "?", "#", "#", "."},
			Broken: []int{1, 1, 3},
		}

		result := input.PossibleBrokenCombinations()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("without gaps", func(t *testing.T) {
		expected := 10
		input := ConditionRecord{
			Raw:    []string{"?", "#", "#", "#", "?", "?", "?", "?", "?", "?", "?", "?"},
			Broken: []int{3, 2, 1},
		}

		result := input.PossibleBrokenCombinations()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
