package part1

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("damaged springs, how original...", func(t *testing.T) {
		expected := 21
		input := []string{
			"???.### 1,1,3",
			".??..??...?##. 1,1,3",
			"?#?#?#?#?#?#?#? 1,3,1,6",
			"????.#...#... 4,1,1",
			"????.######..#####. 1,6,5",
			"?###???????? 3,2,1",
		}

		result := Run(input)

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
