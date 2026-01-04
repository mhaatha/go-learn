package main

import "fmt"

func main() {
	mySlice := []int{8, 5, 3, 9, 4, 1}

	for i := 1; i < len(mySlice); i++ {
		key := mySlice[i]
		j := i - 1

		for j >= 0 && mySlice[j] > key {
			mySlice[j+1] = mySlice[j]
			j--
		}

		mySlice[j+1] = key
	}

	fmt.Println(mySlice)
}
