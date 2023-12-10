package models

import "github.com/samber/lo"

type Vertex struct {
	Index int
	Edges []Edge
}

type Graph []Vertex

func (g Graph) GetByIndex(index int) Vertex {
	element, found := lo.Find(g, func(item Vertex) bool {
		return item.Index == index
	})

	if !found {
		panic("What?!")
	}

	return element
}
