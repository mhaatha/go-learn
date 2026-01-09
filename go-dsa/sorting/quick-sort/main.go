package main

import (
	"fmt"
)

func main() {
	mySlice := []int{8, 5, 3, 9, 4, 1}

	quickSort(mySlice, 0, len(mySlice)-1)

	fmt.Println(mySlice)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		quickSort(arr, low, p-1)

		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return i
}
