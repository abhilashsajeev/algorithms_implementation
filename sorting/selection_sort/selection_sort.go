package main

import (
	"errors"
	"fmt"
	"log/slog"
)

func main() {
	fmt.Println("Selection sort start ")
	arr := []int{5,8,1,4,2,9,1}
	fmt.Printf("%d" ,selectionSort(arr))
}

func selectionSort(arr []int) []int {
	var newArr []int
	copyArr := arr
	for i, _ := range copyArr {
		fmt.Printf("%d", i)
		smallest, err := findSmallest(copyArr)
		if(err != nil){
			slog.Error("Error in panicking")
		}
		newArr = append(newArr, copyArr[smallest])
		copyArr = append(copyArr[:smallest],copyArr[smallest+1:]... )
		
	}
	return newArr



}

func findSmallest(arr []int) (int, error) {
	if(len(arr) == 0) {
		return 0, errors.New("Empty Array")
	}
	smallest := arr[0]
	smallestIndex := 0
	for i, v := range arr {
		if(smallest > v) {
			smallest = v
			smallestIndex = i
		}
	}
	return smallestIndex,nil

}