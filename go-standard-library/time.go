package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.June, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
}
