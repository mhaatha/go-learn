package main

import (
	"fmt"
	"os"
)

const pendingStatus = "pending"

func main() {
	args := os.Args

	switch args[1] {
	case "add":
		if args[2] == "" {
			fmt.Println("Error: task title is required")
			return
		}
		fmt.Println("Task created successfully.")
		fmt.Printf("\nID: 1\nTitle: %s\nStatus: %s\n", args[2], pendingStatus)
	default:
		fmt.Printf("Error: unknown command \"%s\"\n", args[1])
	}
}
