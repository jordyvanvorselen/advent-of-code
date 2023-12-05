package models

import (
	"testing"
)

func TestFindDestination(t *testing.T) {
	t.Run("with non-matching range mapping", func(t *testing.T) {
		source := 79
		expected := source
		input := RangeMapping{
			DestinationRangeStart: 50,
			SourceRangeStart:      98,
			RangeLength:           2,
		}

		result := input.FindDestination(source)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("with matching range mapping", func(t *testing.T) {
		source := 79
		expected := 81
		input := RangeMapping{
			DestinationRangeStart: 52,
			SourceRangeStart:      50,
			RangeLength:           48,
		}

		result := input.FindDestination(source)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
