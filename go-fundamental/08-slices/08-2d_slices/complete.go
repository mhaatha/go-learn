package main

import "fmt"

func createMatric(rows, cols int) [][]int {
	twoD := make([][]int, cols)
	for i := 0; i < cols; i++ {
		twoD[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			twoD[i][j] = i * j
		}
	}
	return twoD
}

// don't edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatric(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
