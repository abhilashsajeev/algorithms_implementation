package main

import "fmt"

func main() {
	fmt.Println(generateSpiralMatrix(4))
}

func generateSpiralMatrix(num int) [][]int {
	startColumn := 0
	startRow := 0
	endColumn := num - 1
	endRow := num - 1

	
    result := make([][]int, num) // Allocate memory for the matrix
    for i := range result {
        result[i] = make([]int, num)
    }


	counter := 1
	
	for startColumn <= endColumn && startRow <= endRow {
		 // Top row
		for i := startColumn; i <= endColumn; i++ {
			result[startRow][i] = counter
			counter++
		}
		startRow++

		// Right most Column
		for i := startRow; i <= endRow; i++ {
			result[i][endColumn] = counter
			counter++
		}
		endColumn--

		// Botton row
		for i := endColumn; i >= startColumn; i-- {
			result[endRow][i] = counter
			counter++
		}
		endRow--
		// Leftmost column
		for i := endRow; i >= startRow; i-- {
			result[i][startColumn] = counter
			counter++
		}
		startColumn++
	}

	return result

}
