package part1

import (
	"2023/day10/part1/models"
	"2023/day10/part1/parsers"
)

func dfs(graph map[int][]models.Edge, visited map[int]bool, current, distance int) (int, int) {
	visited[current] = true
	maxDistance := distance
	farthestVertex := current

	for _, edge := range graph[current] {
		if !visited[edge.To] {
			newDistance := distance + edge.Weight
			vertex, d := dfs(graph, visited, edge.To, newDistance)
			if d > maxDistance {
				maxDistance = d
				farthestVertex = vertex
			}
		}
	}

	return farthestVertex, maxDistance
}

func farthestVertex(graph map[int][]models.Edge, start int) int {
	visited := make(map[int]bool)
	_, maxDistance := dfs(graph, visited, start, 0)
	return maxDistance
}

func Run(input []string) int {
	start, graph := parsers.Parse(input)
	maxDistance := farthestVertex(graph, start)

	return maxDistance/2 + 1
}
