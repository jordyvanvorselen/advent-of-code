package part2

import (
	"2023/day10/part2/models"
	"2023/day10/part2/parsers"
	geo "github.com/kellydunn/golang-geo"
	"slices"
)

// . . . X X X . .
// . X X X . X . .
// . X . . . X . .
// . X . . . X . .
// . X X X X X . .
var loop []int

func dfs(graph models.Graph, visited map[int]bool, current, distance int) (int, int, []int) {
	visited[current] = true
	maxDistance := distance
	farthestVertex := current

	for _, edge := range graph.GetByIndex(current).Edges {
		if !visited[edge.To] {
			loop = append(loop, edge.To)
			newDistance := distance + edge.Weight
			vertex, d, _ := dfs(graph, visited, edge.To, newDistance)
			if d > maxDistance {
				maxDistance = d
				farthestVertex = vertex
			}
		}
	}

	return farthestVertex, maxDistance, loop
}

func coordinatesForIndex(index int, lineLength int) (int, int) {
	var x int
	var y int

	if index == lineLength {
		x = index
	} else {
		x = index % lineLength
	}

	y = index / lineLength

	return x, y
}

func amountOfTilesInLoop(graph models.Graph, start int, lineLength int) int {
	visited := make(map[int]bool)
	_, _, loop := dfs(graph, visited, start, 0)

	var points []*geo.Point

	for _, i := range loop {
		x, y := coordinatesForIndex(i, lineLength)
		points = append(points, geo.NewPoint(float64(x), float64(y)))
	}

	polygon := geo.NewPolygon(points)

	var result int

	for _, g := range graph {
		x, y := coordinatesForIndex(g.Index, lineLength)
		if newPoint := geo.NewPoint(float64(x), float64(y)); !slices.Contains(loop, g.Index) && polygon.Contains(newPoint) {
			result++
		}
	}

	return result
}

func Run(input []string) int {
	start, graph, lineLength := parsers.Parse(input)

	return amountOfTilesInLoop(graph, start, lineLength)
}
