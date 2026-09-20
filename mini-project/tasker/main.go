package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"github.com/mhaatha/go-learn/mini-project/tasker/internal/repository"
	"github.com/mhaatha/go-learn/mini-project/tasker/internal/service"
)

type AddCommand struct {
	Title string `short:"t" long:"title" description:"Task title" required:"true"`
}

type ListCommand struct{}

type DoneCommand struct {
	ID int `short:"i" long:"id" description:"Task ID" required:"true"`
}

type ShowCommand struct {
	ID int `short:"i" long:"id" description:"Task ID" required:"true"`
}

type EditCommand struct {
	ID    int    `short:"i" long:"id" description:"Task ID" required:"true"`
	Title string `short:"t" long:"title" description:"Task title" required:"true"`
}

type DeleteCommand struct {
	ID int `short:"i" long:"id" description:"Task ID" required:"true"`
}

type options struct {
	Add    AddCommand    `command:"add" description:"Create a new task"`
	List   ListCommand   `command:"list" description:"List all tasks"`
	Done   DoneCommand   `command:"done" description:"Mark a task as done"`
	Show   ShowCommand   `command:"show" description:"Show a task"`
	Edit   EditCommand   `command:"edit" description:"Edit a task"`
	Delete DeleteCommand `command:"delete" description:"Delete a task"`
}

func main() {
	var opts options

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "tasker"
	parser.Usage = " "

	_, err := parser.Parse()
	if err != nil {
		return
	}

	taskRepo := repository.NewTaskRepository("database.json")
	taskService := service.NewTaskService(taskRepo)

	switch parser.Command.Active.Name {
	case "add":
		if err := taskService.AddTask(opts.Add.Title); err != nil {
			fmt.Println(err)
		}
	case "list":
		if err := taskService.ListTasks(); err != nil {
			fmt.Println(err)
		}
	case "done":
		if err := taskService.MarkTaskAsDone(opts.Done.ID); err != nil {
			fmt.Println(err)
		}
	case "show":
		if err := taskService.ShowTask(opts.Show.ID); err != nil {
			fmt.Println(err)
		}
	case "edit":
		if err := taskService.EditTask(opts.Edit.ID, opts.Edit.Title); err != nil {
			fmt.Println(err)
		}
	case "delete":
		if err := taskService.DeleteTask(opts.Delete.ID); err != nil {
			fmt.Println(err)
		}
	}
}
