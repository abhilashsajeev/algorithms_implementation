package main

import (
    "testing"
)

func TestGenerateSpiralMatrixEmpty(t *testing.T) {
    expected := [][]int{}
    result := generateSpiralMatrix(0)
    if !compareMatrix(result, expected) {
        t.Errorf("Spiral matrix generation for n=0 failed: expected %v, got %v", expected, result)
    }
}

func TestGenerateSpiralMatrixSingleElement(t *testing.T) {
    expected := [][]int{{1}}
    result := generateSpiralMatrix(1)
    if !compareMatrix(result, expected) {
        t.Errorf("Spiral matrix generation for n=1 failed: expected %v, got %v", expected, result)
    }
}

func TestGenerateSpiralMatrixSimple(t *testing.T) {
    expected := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
    result := generateSpiralMatrix(3)
    if !compareMatrix(result, expected) {
        t.Errorf("Spiral matrix generation for n=3 failed: expected %v, got %v", expected, result)
    }
}

func TestGenerateSpiralMatrixLarger(t *testing.T) {
    expected := [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}}
    result := generateSpiralMatrix(5)
    if !compareMatrix(result, expected) {
        t.Errorf("Spiral matrix generation for n=5 failed: expected %v, got %v", expected, result)
    }
}

func compareMatrix(a [][]int, b [][]int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if !compareIntSlices(a[i], b[i]) {
            return false
        }
    }
    return true
}

// Helper function for comparing integer slices
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
