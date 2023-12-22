// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, 世界")
	unsort := []int{4, 8, 1, 2, 5}
	fmt.Println(quickSort(unsort))
}

// quick sort with random element as pivot
func quickSortRand(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivotIndex := rand.Intn(len(arr))
		pivot := arr[pivotIndex]

		// Move the pivot to the end
		arr[pivotIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotIndex]
		var less []int
		var greater []int
		for _, v := range arr[1:] {
			if v <= pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
		finalval := append(append(quickSort(less), pivot), quickSort(greater)...)
		return finalval
	}

}

// quick sort with first element as pivot
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less []int
		var greater []int
		for _, v := range arr[1:] {
			if v <= pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
		finalval := append(append(quickSort(less), pivot), quickSort(greater)...)
		return finalval
	}

}
