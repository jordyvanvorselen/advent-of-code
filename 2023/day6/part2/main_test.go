package part2

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 71503
	input := []string{
		"Time:      71530",
		"Distance:  940200",
	}

	result := Run(input)

	if result != expected {
		t.Errorf("got %d but expected %d", result, expected)
	}
}
