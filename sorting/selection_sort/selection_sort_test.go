package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 8, 1, 2, 5}, []int{1, 2, 4, 5, 8}},
		{[]int{3, 7, 1, 4, 6, 8, 2, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		result := selectionSort(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("For input %v, expected %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
