// https://www.codewars.com/kata/52b757663a95b11b3d00062d/train/go

package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	newStr := ""
	upper := 1
	lower := 0

	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " && upper > lower {
			newStr += string(str[i])
			continue
		}
		if upper > lower {
			newStr += strings.ToUpper(string(str[i]))
			lower++
		} else {
			newStr += strings.ToLower(string(str[i]))
			upper++
		}
	}
	return newStr
}

func main() {
	fmt.Println(toWeirdCase("String"))            // StRiNg
	fmt.Println(toWeirdCase("Weird string case")) // WeIrD StRiNg CaSe
}
