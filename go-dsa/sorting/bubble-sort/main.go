package main

import "fmt"

func main() {
	mySlice := []int{64, 34, 25, 12, 22, 11, 90}

	for i := 0; i < len(mySlice)-1; i++ {
		swapped := false
		for j := 0; j < len(mySlice)-i-1; j++ {
			if mySlice[j] > mySlice[j+1] {
				mySlice[j], mySlice[j+1] = mySlice[j+1], mySlice[j]
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}

	fmt.Println(mySlice)
}
