package main

import "fmt"

func main() {
	var name string

	name = "Jokowi"
	fmt.Println(name)

	name = "Joestar"
	fmt.Println(name)

	var langsung = "Variable string langsung" // Inisialisasi data type string tidak wajib lagi karena variable name sudah diinisialisasi menjadi string, data type pada sebuah variable tidak dapat diubah.
	fmt.Println(langsung)

	noVar := "Variable tanpa var" // Keyword var tidak wajib asal variable tersebut diinisialisasi datanya dengan cara menambahkan := tapi ini dilakukan hanya pada deklarasi pertama saja
	fmt.Println(noVar)

	noVar = "Apa itu var?"
	fmt.Println(noVar)

	// Multi variable
	var (
		firstName  = "Jokowi"
		middleName = "Joestar"
		lastName   = "Dodo"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
