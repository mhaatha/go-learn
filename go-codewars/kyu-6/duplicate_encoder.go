package main

import (
	"fmt"
	"strings"
)

func duplicateEncode(word string) string {
	result := ""
	letter := make(map[string]int)
	for _, w := range word {
		currentLetter := letter[strings.ToLower(string(w))]
		if currentLetter == 0 {
			letter[strings.ToLower(string(w))] = 1
		} else {
			letter[strings.ToLower(string(w))]++
		}
	}

	for _, w := range word {
		if letter[strings.ToLower(string(w))] > 1 {
			result += ")"
		} else {
			result += "("
		}
	}

	return result
}

func main() {
	fmt.Println(duplicateEncode("din"))     // (((
	fmt.Println(duplicateEncode("recede"))  // ()()()
	fmt.Println(duplicateEncode("Success")) // )())())
	fmt.Println(duplicateEncode("(( @"))    // ))((
}
