package main

import (
	"fmt"
)

func main() {
	mySlice := []int{8, 5, 3, 9, 4, 1}

	mySortedSlice := mergeSort(mySlice)

	fmt.Println(mySortedSlice)
}

func mergeSort(items []int) []int {
	// Base case
	if len(items) < 2 {
		return items
	}

	// Divide the slice into two
	midValue := len(items) / 2

	leftHalf := items[:midValue]
	rightHalf := items[midValue:]

	sortedLeft := mergeSort(leftHalf)   // Pause until finish...
	sortedRight := mergeSort(rightHalf) // Pause until finish...

	return merge(sortedLeft, sortedRight)
}

func merge(left, right []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		// Bandingkan left[i] dengan right[j]
		// Angka yang lebih kecil, akan diappend ke final setelah itu indexnya++
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}

	// Cleanup
	final = append(final, left[i:]...)
	final = append(final, right[j:]...)

	return final
}
