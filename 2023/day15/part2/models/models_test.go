package models

import "testing"

func TestInitItem(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expected := 0
		input := InitItem{Label: "rn", Op: ADD, FocalLength: 1}

		result := input.Hash()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		expected := 0
		input := InitItem{Label: "cm", Op: REMOVE}

		result := input.Hash()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		expected := 3
		input := InitItem{Label: "pc", Op: ADD, FocalLength: 4}

		result := input.Hash()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
