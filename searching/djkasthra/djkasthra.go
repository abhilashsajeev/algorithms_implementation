package main

import (
	"fmt"
	"math"
)

func dijkstra(graph [][]int, startNode int) []int {
	numNodes := len(graph)
	distances := make([]int, numNodes)
	visited := make([]bool, numNodes)

	// Initialize distances with infinity and mark startNode distance as 0
	for i := 0; i < numNodes; i++ {
		distances[i] = math.MaxInt32
	}
	distances[startNode] = 0

	for count := 0; count < numNodes-1; count++ {
		// Find the minimum distance vertex from the set of vertices not yet processed
		u := minDistance(distances, visited)

		// Mark the picked vertex as processed
		visited[u] = true

		// Update the distance value of the adjacent vertices of the picked vertex
		for v := 0; v < numNodes; v++ {
			if !visited[v] && graph[u][v] != 0 && distances[u] != math.MaxInt32 &&
				distances[u]+graph[u][v] < distances[v] {
				distances[v] = distances[u] + graph[u][v]
			}
		}
	}

	return distances
}

func minDistance(distances []int, visited []bool) int {
	min := math.MaxInt32
	minIndex := -1

	for v := 0; v < len(distances); v++ {
		if !visited[v] && distances[v] <= min {
			min = distances[v]
			minIndex = v
		}
	}

	return minIndex
}

func main() {
	// Example graph represented as an adjacency matrix
	graph := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}

	startNode := 0
	distances := dijkstra(graph, startNode)

	fmt.Printf("Shortest distances from node %d:\n", startNode)
	for i, distance := range distances {
		fmt.Printf("Node %d: %d\n", i, distance)
	}
}
