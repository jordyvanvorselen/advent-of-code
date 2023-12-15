package part1

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("oooooooooooooh hashes... shouldn't be too hard.", func(t *testing.T) {
		expected := 1320
		input := []string{
			"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
