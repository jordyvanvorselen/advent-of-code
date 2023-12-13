package models

import "testing"

func TestMirrorList(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expected := 2
		input := MirrorList{
			Rows: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
		}

		result, found, org := input.FindPerfectRowReflectionIndex()

		if !found {
			t.Error("expected index to be found!")
		}

		if org {
			t.Error("expected a new reflection to be found!")
		}

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		expected := 0
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

		result, found, org := input.FindPerfectRowReflectionIndex()

		if !found {
			t.Error("expected index to be found!")
		}

		if org {
			t.Error("expected a new reflection to be found!")
		}

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		expected := 10
		input := MirrorList{
			Columns: []string{
				"######.#.##.#..",
				"#.###.#.##.#.##",
				".#...#.####.##.",
				"#####.###.....#",
				".#.##..#..#..#.",
				".##.#..##....##",
				"..#.#..#..###..",
				"..#.#..#..###..",
				".##.#..#.....##",
				".##.#..#.....##",
				"..#.#..#..###..",
				"..#.#..#..###..",
				".##.#..##....##",
			},
		}

		result, found, org := input.FindPerfectColumnReflectionIndex()

		if !found {
			t.Error("expected index to be found!")
		}

		if org {
			t.Error("expected a new reflection to be found!")
		}

		if result != expected {
			t.Errorf("got %d but expected %d", result, expected)
		}
	})
}
