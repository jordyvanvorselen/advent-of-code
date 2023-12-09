package models

import (
	"testing"
)

func TestSequence_Prediction(t *testing.T) {
	t.Run("sequence 1", func(t *testing.T) {
		expected := -3
		input := Sequence{0, 3, 6, 9, 12, 15}

		result := input.Prediction()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("sequence 2", func(t *testing.T) {
		expected := 0
		input := Sequence{1, 3, 6, 10, 15, 21}

		result := input.Prediction()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("sequence 3", func(t *testing.T) {
		expected := 5
		input := Sequence{10, 13, 16, 21, 30, 45}

		result := input.Prediction()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
