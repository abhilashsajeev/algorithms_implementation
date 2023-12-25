package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{4, 8, 1, 2, 5}
	expected := []int{1, 2, 4, 5, 8}

	result := mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
