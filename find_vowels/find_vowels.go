// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
)

func main() {
	test := "This is test string"
	fmt.Printf("Count by normal fn %d", countVowels(test))
	fmt.Printf("\n Count by Regx fn %d", countByRegx(test))

}

func countByRegx(str string) int {
	comp := regexp.MustCompile("(?i)[aeiou]")

	strLen := comp.FindAllString(str, -1)
	return len(strLen)

}

// Normal way by looping and checking count using this
func countVowels(str string) int {
	checker := []string{"a", "i", "e", "o", "u"}
	count := 0
	for _, c := range str {
		var lower string
		if c >= 'A' && c <= 'Z' {
			lower = string(c + 32)
		} else {
			lower = string(c)
		}
		if Contains(checker, string(lower)) {
			count++
		}
	}
	return count
}

// Contains checks if a slice contains an element, using generics
func Contains[T comparable](slice []T, elem T) bool {
	for _, t := range slice {
		if t == elem {
			return true
		}
	}
	return false
}
