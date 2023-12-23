package main

import (
    "errors"
    "fmt"
)

func main() {
    fmt.Println("Selection sort start")
    arr := []int{5, 8, 1, 4, 2, 9, 1}
    sortedArr := selectionSort(arr)
    fmt.Println(sortedArr) // Print the sorted array
}

func selectionSort(arr []int) []int {
    newArr := make([]int, 0, len(arr)) // Pre-allocate newArr for efficiency
    for len(arr) > 0 {
        smallest, err := findSmallest(arr)
        if err != nil {
            panic(err) // Handle the error appropriately
        }
        newArr = append(newArr, arr[smallest])
        arr = append(arr[:smallest], arr[smallest+1:]...)
    }
    return newArr
}

func findSmallest(arr []int) (int, error) {
    if len(arr) == 0 {
        return 0, errors.New("Empty Array")
    }
    smallestIndex := 0
    for i, v := range arr {
        if v < arr[smallestIndex] {
            smallestIndex = i
        }
    }
    return smallestIndex, nil
}
