package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Byte Conversion
	var name = "jokowi"
	var j = name[0]       // Output: j in Byte
	var jWord = string(j) // Output: j

	fmt.Println(j)
	fmt.Println(jWord)
}
