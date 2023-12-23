package main

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	testCases := []struct {
		graph     [][]int
		startNode int
		expected  []int
	}{
		{
			graph: [][]int{
				{0, 4, 0, 0, 0, 0, 0, 8, 0},
				{4, 0, 8, 0, 0, 0, 0, 11, 0},
				{0, 8, 0, 7, 0, 4, 0, 0, 2},
				{0, 0, 7, 0, 9, 14, 0, 0, 0},
				{0, 0, 0, 9, 0, 10, 0, 0, 0},
				{0, 0, 4, 14, 10, 0, 2, 0, 0},
				{0, 0, 0, 0, 0, 2, 0, 1, 6},
				{8, 11, 0, 0, 0, 0, 1, 0, 7},
				{0, 0, 2, 0, 0, 0, 6, 7, 0},
			},
			startNode: 0,
			expected:  []int{0, 4, 12, 19, 21, 11, 9, 8, 14},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		result := dijkstra(tc.graph, tc.startNode)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("For startNode %d, expected %v, but got %v", tc.startNode, tc.expected, result)
		}
	}
}
