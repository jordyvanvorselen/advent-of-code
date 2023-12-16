package models

import "fmt"

type Tile struct {
	X    int
	Y    int
	Type TileType
}

type Grid map[string]Tile

type TileType string

const (
	EMPTY               TileType = "."
	VERTICAL_SPLITTER   TileType = "|"
	HORIZONTAL_SPLITTER TileType = "-"
	MIRROR_SLASH        TileType = "/"
	MIRROR_BASHSLASH    TileType = `\`
)

type Direction string

const (
	UP    Direction = "N"
	LEFT  Direction = "W"
	DOWN  Direction = "S"
	RIGHT Direction = "E"
)

type Step struct {
	Tile      Tile
	Direction Direction
}

func (g Grid) GetNext(current Step) []Step {
	var result []Step

	x, y := current.getNewCoords()

	if g.outOfBounds(x, y) {
		return []Step{}
	}

	new := g[coords(x, y)]

	switch new.Type {
	case EMPTY:
		return []Step{{Tile: new, Direction: current.Direction}}
	case VERTICAL_SPLITTER:
		if current.Direction == LEFT || current.Direction == RIGHT {
			return []Step{
				{Tile: new, Direction: UP},
				{Tile: new, Direction: DOWN},
			}
		}

		return []Step{{Tile: new, Direction: current.Direction}}
	case HORIZONTAL_SPLITTER:
		if current.Direction == UP || current.Direction == DOWN {
			return []Step{
				{Tile: new, Direction: LEFT},
				{Tile: new, Direction: RIGHT},
			}
		}

		return []Step{{Tile: new, Direction: current.Direction}}
	case MIRROR_SLASH: //        /
		if current.Direction == UP {
			return []Step{{Tile: new, Direction: RIGHT}}
		}

		if current.Direction == DOWN {
			return []Step{{Tile: new, Direction: LEFT}}
		}

		if current.Direction == RIGHT {
			return []Step{{Tile: new, Direction: UP}}
		}

		if current.Direction == LEFT {
			return []Step{{Tile: new, Direction: DOWN}}
		}
	case MIRROR_BASHSLASH:
		if current.Direction == UP {
			return []Step{{Tile: new, Direction: LEFT}}
		}

		if current.Direction == DOWN {
			return []Step{{Tile: new, Direction: RIGHT}}
		}

		if current.Direction == RIGHT {
			return []Step{{Tile: new, Direction: DOWN}}
		}

		if current.Direction == LEFT {
			return []Step{{Tile: new, Direction: UP}}
		}
	}

	return result
}

func (g Grid) outOfBounds(x int, y int) bool {
	_, ok := g[coords(x, y)]

	return !ok
}

func (s Step) getNewCoords() (int, int) {
	if s.Direction == UP {
		return s.Tile.X, s.Tile.Y - 1
	}

	if s.Direction == DOWN {
		return s.Tile.X, s.Tile.Y + 1
	}

	if s.Direction == LEFT {
		return s.Tile.X - 1, s.Tile.Y
	}

	if s.Direction == RIGHT {
		return s.Tile.X + 1, s.Tile.Y
	}

	panic("Cannot determine new coords.")
}

func coords(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
