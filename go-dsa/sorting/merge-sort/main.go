package main

import (
	"math"
)

func main() {

}

func mergeSort(items []int) []int {
	// items = [38, 27, 43, 3]

	// Base case
	if len(items) < 2 {
		return items
	}

	// Divide the slice into two
	midValue := int(math.Floor(float64(len(items) / 2)))

	leftHalf := items[:midValue]  // [38, 27]
	rightHalf := items[midValue:] // [43, 3]

	sortedLeft := mergeSort(leftHalf)
	sortedRight := mergeSort(rightHalf)

	// Waiting...

	return merge(sortedLeft, sortedRight)
}

func merge(left, right []int) []int {
	final := []int{}

	return final
}
