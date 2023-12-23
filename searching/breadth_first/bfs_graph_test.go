package breadth_first

import (
    "testing"
)

func TestBFSEmptyGraph(t *testing.T) {
    graph := Graph{nodes: make(map[string][]string)}
    expected := []string{}
    result := graph.BFS("A")
    if !compareStringSlices(result, expected) {
        t.Errorf("BFS on empty graph failed: expected %v, got %v", expected, result)
    }
}

func TestBFSSingleNodeGraph(t *testing.T) {
    graph := Graph{nodes: map[string][]string{"A": {}}}
    expected := []string{"A"}
    result := graph.BFS("A")
    if !compareStringSlices(result, expected) {
        t.Errorf("BFS on single node graph failed: expected %v, got %v", expected, result)
    }
}

func TestBFSSimpleGraph(t *testing.T) {
    graph := Graph{nodes: map[string][]string{"A": {"B", "C"}, "B": {"D"}, "C": {"E"}}}
    expected := []string{"A", "B", "C", "D", "E"}
    result := graph.BFS("A")
    if !compareStringSlices(result, expected) {
        t.Errorf("BFS on simple graph failed: expected %v, got %v", expected, result)
    }
}

func TestBFSGraphWithDisconnectedNodes(t *testing.T) {
    graph := Graph{nodes: map[string][]string{"A": {"B", "C"}, "X": {"Y"}}}
    expected := []string{"A", "B", "C"}
    result := graph.BFS("A")
    if !compareStringSlices(result, expected) {
        t.Errorf("BFS on graph with disconnected nodes failed: expected %v, got %v", expected, result)
    }
}

func compareStringSlices(a []string, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
