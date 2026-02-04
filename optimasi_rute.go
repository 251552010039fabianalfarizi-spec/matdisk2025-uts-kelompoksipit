package main

import (
	"math"
)

// Struktur Graph
type Edge struct {
	to     string
	weight int
}

type Graph map[string][]Edge

// Dijkstra
func Dijkstra(graph Graph, start string) map[string]int {
	dist := make(map[string]int)
	visited := make(map[string]bool)

	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	for {
		current := ""
		minDist := math.MaxInt32

		for node, d := range dist {
			if !visited[node] && d < minDist {
				minDist = d
				current = node
			}
		}

		if current == "" {
			break
		}

		visited[current] = true

		for _, edge := range graph[current] {
			if dist[current]+edge.weight < dist[edge.to] {
				dist[edge.to] = dist[current] + edge.weight
			}
		}
	}

	return dist
}
