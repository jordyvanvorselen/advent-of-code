package models

import "testing"

func TestInitItem(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expected := 30
		input := InitItem{Raw: "rn=1"}

		result := input.Hash()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		expected := 253
		input := InitItem{Raw: "cm-"}

		result := input.Hash()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		expected := 97
		input := InitItem{Raw: "qp=3"}

		result := input.Hash()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		expected := 47
		input := InitItem{Raw: "cm=2"}

		result := input.Hash()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
