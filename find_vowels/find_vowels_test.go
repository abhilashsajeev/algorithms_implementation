package main

import (
	"testing"
)

func TestCountByRegx(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"This is test string", 4},
		{"AEIOUaeiou", 10}, // Check for case-insensitive matching
		{"12345", 0},       // No vowels
	}

	for _, tc := range testCases {
		result := countByRegx(tc.input)
		if result != tc.expected {
			t.Errorf("For input '%s', expected %d, but got %d", tc.input, tc.expected, result)
		}
	}
}

func TestCountVowels(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"This is test string", 4},
		{"AEIOUaeiou", 10}, // Check for case-insensitive matching
		{"12345", 0},       // No vowels
	}

	for _, tc := range testCases {
		result := countVowels(tc.input)
		if result != tc.expected {
			t.Errorf("For input '%s', expected %d, but got %d", tc.input, tc.expected, result)
		}
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		slice    []string
		elem     string
		expected bool
	}{
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"x", "y", "z"}, "a", false},
		{[]string{"hello", "world"}, "world", true},
	}

	for _, tc := range testCases {
		result := Contains(tc.slice, tc.elem)
		if result != tc.expected {
			t.Errorf("For slice %v and elem %s, expected %t, but got %t", tc.slice, tc.elem, tc.expected, result)
		}
	}
}
