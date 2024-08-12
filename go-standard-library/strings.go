package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hafidz Athaya", "fidz"))
	fmt.Println(strings.Split(" Hafidz Athaya ", " "))
	fmt.Println(strings.ToLower("Hafidz Athaya"))
	fmt.Println(strings.ToUpper("Hafidz Athaya"))
	fmt.Println(strings.Trim("   Hafidz Athaya", " "))
	fmt.Println(strings.ReplaceAll("oink oink oink", "ink", "moo"))
}
