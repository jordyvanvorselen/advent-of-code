package models

import "testing"

func TestMirrorList(t *testing.T) {
	t.Run("vertical", func(t *testing.T) {
		expected := 4
		input := MirrorList{
			Columns: []string{
				"#.##..#",
				"..##...",
				"##..###",
				"#....#.",
				".#..#.#",
				".#..#.#",
				"#....#.",
				"##..###",
				"..##...",
			},
		}

		result, found := input.FindPerfectColumnReflectionIndex()

		if !found {
			t.Error("expected index to be found!")
		}

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("horizontal", func(t *testing.T) {
		expected := 3
		input := MirrorList{
			Rows: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
		}

		result, found := input.FindPerfectRowReflectionIndex()

		if !found {
			t.Error("expected index to be found!")
		}

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
