package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
	unsort := []int{4, 8, 1, 2, 5}
	fmt.Println(mergeSort(unsort))
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	center := len(arr) / 2
	left := arr[:center ]
	right := arr[center:]
	return merge(mergeSort(left), mergeSort(right))

}

func merge(left []int, right []int) []int {
	var results []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			results = append(results, left[0])
			left = left[1:]
		} else {
			results = append(results, right[0])
			right = right[1:]
		}

	}
	return append(append(results, left...), right...)

}
