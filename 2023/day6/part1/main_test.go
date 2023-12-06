package part1

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 288
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	result := Run(input)

	if result != expected {
		t.Errorf("got %d but expected %d", result, expected)
	}
}
