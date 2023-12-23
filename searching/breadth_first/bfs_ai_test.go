package breadth_first

import (
    "testing"
)

func TestBFSEmptyTree(t *testing.T) {
    tree := &Tree{}
    expected := []int{}
    result := BFS(tree)
    if !compareIntSlices(result, expected) {
        t.Errorf("BFS on empty tree failed: expected %v, got %v", expected, result)
    }
}

func TestBFSSingleNodeAI(t *testing.T) {
    tree := &Tree{Value: 10}
    expected := []int{10}
    result := BFS(tree)
    if !compareIntSlices(result, expected) {
        t.Errorf("BFS on single node tree failed: expected %v, got %v", expected, result)
    }
}

func TestBFSSimpleTree(t *testing.T) {
    tree := &Tree{Value: 5, Left: &Tree{Value: 3}, Right: &Tree{Value: 7}}
    expected := []int{5, 3, 7}
    result := BFS(tree)
    if !compareIntSlices(result, expected) {
        t.Errorf("BFS on simple tree failed: expected %v, got %v", expected, result)
    }
}

func TestBFSComplexTree(t *testing.T) {
    tree := &Tree{
        Value: 1,
        Left: &Tree{Value: 2, Left: &Tree{Value: 4}, Right: &Tree{Value: 5}},
        Right: &Tree{Value: 3, Left: &Tree{Value: 6}, Right: &Tree{Value: 7}},
    }
    expected := []int{1, 2, 3, 4, 5, 6, 7}
    result := BFS(tree)
    if !compareIntSlices(result, expected) {
        t.Errorf("BFS on complex tree failed: expected %v, got %v", expected, result)
    }
}

func compareIntSlices(a []int, b []int) bool {
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
