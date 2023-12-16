package models

import (
	"reflect"
	"testing"
)

func TestGrid_GetNext(t *testing.T) {
	grid := Grid{
		"0,0": Tile{X: 0, Y: 0, Type: EMPTY},
		"1,0": Tile{X: 1, Y: 0, Type: HORIZONTAL_SPLITTER},
		"2,0": Tile{X: 2, Y: 0, Type: MIRROR_BASHSLASH},
		"0,1": Tile{X: 0, Y: 1, Type: MIRROR_SLASH},
		"1,1": Tile{X: 1, Y: 1, Type: VERTICAL_SPLITTER},
		"2,1": Tile{X: 2, Y: 1, Type: EMPTY},
		"0,2": Tile{X: 0, Y: 2, Type: MIRROR_BASHSLASH},
		"1,2": Tile{X: 1, Y: 2, Type: EMPTY},
		"2,2": Tile{X: 2, Y: 2, Type: MIRROR_SLASH},
	}

	t.Run("when hitting empty space", func(t *testing.T) {
		step := Step{
			Tile:      grid["0,0"],
			Direction: RIGHT,
		}

		expected := []Step{
			{
				Tile:      grid["1,0"],
				Direction: RIGHT,
			},
		}

		result := grid.GetNext(step)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %+v but expected %+v", result, expected)
		}
	})

	t.Run("when hitting backslash mirror", func(t *testing.T) {
		step := Step{
			Tile:      grid["1,0"],
			Direction: RIGHT,
		}

		expected := []Step{
			{
				Tile:      grid["2,0"],
				Direction: DOWN,
			},
		}

		result := grid.GetNext(step)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %+v but expected %+v", result, expected)
		}
	})

	t.Run("when hitting slash mirror", func(t *testing.T) {
		step := Step{
			Tile:      grid["2,1"],
			Direction: DOWN,
		}

		expected := []Step{
			{
				Tile:      grid["2,2"],
				Direction: LEFT,
			},
		}

		result := grid.GetNext(step)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %+v but expected %+v", result, expected)
		}
	})

	t.Run("when hitting vertical splitter", func(t *testing.T) {
		step := Step{
			Tile:      grid["0,1"],
			Direction: RIGHT,
		}

		expected := []Step{
			{
				Tile:      grid["1,1"],
				Direction: UP,
			},
			{
				Tile:      grid["1,1"],
				Direction: DOWN,
			},
		}

		result := grid.GetNext(step)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %+v but expected %+v", result, expected)
		}
	})

	t.Run("when hitting horizontal splitter", func(t *testing.T) {
		step := Step{
			Tile:      grid["1,1"],
			Direction: UP,
		}

		expected := []Step{
			{
				Tile:      grid["1,0"],
				Direction: LEFT,
			},
			{
				Tile:      grid["1,0"],
				Direction: RIGHT,
			},
		}

		result := grid.GetNext(step)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %+v but expected %+v", result, expected)
		}
	})
}
