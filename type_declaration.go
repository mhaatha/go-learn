package main

import "fmt"

func main() {
	type NoKTP string

	var ktpJoko NoKTP = "111111111"
	var contoh string = "222222222"

	// Data Type Conversion
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpJoko)
	fmt.Println(contohKtp)
}
