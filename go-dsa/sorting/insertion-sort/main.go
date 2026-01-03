package main

import "fmt"

func main() {
	mySlice := []int{2, 8, 5, 3, 9, 4, 1}

	for i := 1; i < len(mySlice); i++ {
		key := mySlice[i]
		lastIndexOfJ := i - 1

		for j := lastIndexOfJ; j >= 0; j-- {
			if key < mySlice[j] {
				mySlice[j+1] = mySlice[j]
			}
			if key > mySlice[j] {
				break
			}
		}

		mySlice[lastIndexOfJ+1] = key
	}

	fmt.Println(mySlice)

	// PR:
	// 1. Kalau key >= mySlice[j]
	//    maka:
	//    break
	//
	// 2. Swap dilakukan di akhir
}
