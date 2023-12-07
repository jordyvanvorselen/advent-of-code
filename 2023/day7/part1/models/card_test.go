package models

import (
	"testing"
)

func TestCard_Strength(t *testing.T) {
	t.Run("for a number", func(t *testing.T) {
		expected := 3
		input := Card{"5"}

		result := input.Strength()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("for an image card", func(t *testing.T) {
		expected := 11
		input := Card{"K"}

		result := input.Strength()

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
