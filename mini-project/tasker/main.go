package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const pendingStatus = "pending"

type Task struct {
	ID        int
	Title     string
	Status    string
	CreatedAt string
	UpdatedAt string
}

type Database struct {
	NextID int
	Tasks  []Task
}

func main() {
	args := os.Args

	switch args[1] {
	case "add":
		if args[2] == "" {
			fmt.Println("Error: task title is required")
			return
		}

		file, err := os.OpenFile("database.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		dec := jsontext.NewDecoder(file)
		var data Database
		for {
			if err := json.UnmarshalDecode(dec, &data); err != nil {
				if err == io.EOF {
					break
				} else {
					log.Fatal(err)
				}
			}
			data.Tasks = append(data.Tasks, Task{
				ID:        data.NextID,
				Title:     args[2],
				Status:    pendingStatus,
				CreatedAt: time.Now().Format(time.DateTime),
				UpdatedAt: time.Now().Format(time.DateTime),
			})
		}

		currentID := data.NextID
		data.NextID++

		jsonData, err := json.Marshal(data, jsontext.Multiline(true))
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile("database.json", jsonData, 0600)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Task created successfully.")
		fmt.Printf("\nID: %d\nTitle: %s\nStatus: %s\n", currentID, args[2], pendingStatus)
	case "list":
	case "show":
	case "edit":
	case "done":
	case "delete":
	case "help":
	default:
		fmt.Printf("Error: unknown command \"%s\"\n", args[1])
	}
}
