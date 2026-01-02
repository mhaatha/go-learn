package main

import "fmt"

func main() {
	mySlice := []int{2, 8, 5, 3, 9, 4, 1}

	for i := 0; i < len(mySlice)-1; i++ {
		currentMinimum := i

		for j := i + 1; j < len(mySlice); j++ {
			if mySlice[currentMinimum] > mySlice[j] {
				currentMinimum = j
			}
		}

		mySlice[i], mySlice[currentMinimum] = mySlice[currentMinimum], mySlice[i]
	}

	fmt.Println(mySlice)
}
