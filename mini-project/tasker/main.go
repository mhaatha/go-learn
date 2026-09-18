package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"
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
		file, err := os.Open("database.json")
		if err != nil {
			fmt.Println("No tasks found.")
			return
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
		}

		// Empty tasks validation
		if len(data.Tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		// Tabwriter initialization
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

		// Print header
		fmt.Fprintln(w, "ID\tSTATUS\tTITLE")
		fmt.Fprintln(w, "--\t------\t-----")

		// Print tasks
		for _, task := range data.Tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, task.Status, task.Title)
		}

		// Flush the tabwriter
		w.Flush()
	case "show":
	case "edit":
	case "done":
	case "delete":
	case "help":
	default:
		fmt.Printf("Error: unknown command \"%s\"\n", args[1])
	}
}
