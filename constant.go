package main

import "fmt"

func main() {
	const firstName string = "Jokowi" // string tidak wajib
	const lastName = "Joestar"        // unused constant tidak error, kecuali unused var baru error

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Multiple constant
	const (
		first  = "Satu"
		second = "Dua"
	)

	fmt.Println(first)
	fmt.Println(second)
}
