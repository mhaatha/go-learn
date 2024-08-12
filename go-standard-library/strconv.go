package main

import (
	"fmt"
	"strconv"
)

func main() {
	if resultBool, err := strconv.ParseBool("BENAR"); err == nil {
		fmt.Println(resultBool)
	} else {
		fmt.Println("Error:", err)
	}

	if resultInt, err := strconv.Atoi("SALAH"); err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error:", err)
	}

	binary := strconv.FormatInt(1202, 2)
	fmt.Println(binary)

	stringInt := strconv.Itoa(999)
	fmt.Println(stringInt)
}
