// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	unsort := []int{4, 8, 1, 2, 5}
	fmt.Println(bubbleSort(unsort))
}
func bubbleSort(arr []int) []int {
	var swapped bool
	var i, j int
	for i = 0; i < len(arr); i++ {
		swapped = false
		for j = 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	return arr

}
