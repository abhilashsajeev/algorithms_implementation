package main

import (
	"fmt"
)

// Graph structure
type Graph struct {
	nodes map[string][]string
}

// Add a node to the graph
func (g *Graph) addNode(node string) {
	g.nodes[node] = []string{}
}

// Add an edge between two nodes
func (g *Graph) addEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to)
}

// Breadth First Search function
func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !visited[node] {
			visited[node] = true
			fmt.Println(node)

			for _, neighbor := range g.nodes[node] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	graph := Graph{nodes: make(map[string][]string)}
	// Create a sample graph
	graph.addEdge("A", "B")
	graph.addEdge("A", "C")
	graph.addEdge("B", "D")
	graph.addEdge("C", "E")

	graph.BFS("A") // Output: A B C D E
}
