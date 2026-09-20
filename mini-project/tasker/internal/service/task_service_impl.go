package service

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"text/tabwriter"
	"time"

	"github.com/mhaatha/go-learn/mini-project/tasker/internal/model"
	"github.com/mhaatha/go-learn/mini-project/tasker/internal/repository"
)

func NewTaskService(repo repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		Repository: repo,
	}
}

type TaskServiceImpl struct {
	Repository repository.TaskRepository
}

func (a *TaskServiceImpl) AddTask(title string) error {
	// Call the repository to get the current task list
	data, err := a.Repository.ReadCurrentDatabase()
	if err != nil {
		return err
	}

	// Append the new task to the current tasks slice
	data.Tasks = append(data.Tasks, model.Task{
		ID:        data.NextID,
		Title:     title,
		Status:    model.PendingStatus,
		CreatedAt: time.Now().Format(time.DateTime),
		UpdatedAt: time.Now().Format(time.DateTime),
	})

	// Assign NextID to the current ID and increment NextID for the next task
	currentID := data.NextID
	data.NextID++

	// Call the repository to save the task
	if err = a.Repository.OverwriteTask(data); err != nil {
		return err
	}

	// Print the task details
	fmt.Println("Task created successfully.")
	fmt.Printf("\nID: %d\nTitle: %s\nStatus: %s\n", currentID, title, model.PendingStatus)
	return nil
}

func (a *TaskServiceImpl) ListTasks() error {
	// Call the repository to get the current task list
	data, err := a.Repository.ReadCurrentDatabase()
	if err != nil {
		return err
	}

	// Empty tasks validation
	if len(data.Tasks) == 0 {
		return errors.New("No tasks found.")
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
	return nil
}

func (a *TaskServiceImpl) MarkTaskAsDone(ID int) error {
	// Call the repository to get the current task list
	data, err := a.Repository.ReadCurrentDatabase()
	if err != nil {
		return err
	}

	isFound := false
	// Change status to "completed" if task is found
	for i := range data.Tasks {
		if data.Tasks[i].ID == ID {
			// Check if task is already completed
			if data.Tasks[i].Status == "completed" {
				return fmt.Errorf("Error: task %d is already completed.", ID)
			}

			// Update status and timestamp
			data.Tasks[i].Status = "completed"
			data.Tasks[i].UpdatedAt = time.Now().Format(time.DateTime)
			isFound = true
		}
	}

	// Task not found validation
	if !isFound {
		return fmt.Errorf("Error: task %d not found.", ID)
	}

	// Call the repository to save the new task status
	if err = a.Repository.OverwriteTask(data); err != nil {
		return err
	}

	fmt.Printf("Task %d marked as completed.\n", ID)
	return nil
}

func (a *TaskServiceImpl) ShowTask(ID int) error {
	// Call the repository to get the current task list
	data, err := a.Repository.ReadCurrentDatabase()
	if err != nil {
		return err
	}

	isFound := false
	// Find the task by ID and print its details
	for i := range data.Tasks {
		if data.Tasks[i].ID == ID {
			fmt.Printf("ID: %d\nTitle: %s\nStatus: %s\nCreated: %s\nUpdated: %s\n", data.Tasks[i].ID, data.Tasks[i].Title, data.Tasks[i].Status, data.Tasks[i].CreatedAt, data.Tasks[i].UpdatedAt)
			isFound = true
		}
	}

	// Task not found validation
	if !isFound {
		return fmt.Errorf("Error: task with ID %d not found", ID)
	}
	return nil
}

func (a *TaskServiceImpl) EditTask(ID int, title string) error {
	// Call the repository to get the current task list
	data, err := a.Repository.ReadCurrentDatabase()
	if err != nil {
		return err
	}

	isFound := false
	// Find the task by ID and update its title and updatedAt fields
	for i := range data.Tasks {
		if data.Tasks[i].ID == ID {
			data.Tasks[i].Title = title
			data.Tasks[i].UpdatedAt = time.Now().Format(time.DateTime)
			isFound = true
		}
	}

	// Task not found validation
	if !isFound {
		return fmt.Errorf("Error: task %d not found.", ID)
	}

	// Overwrite updated tasks back to file
	if err = a.Repository.OverwriteTask(data); err != nil {
		return err
	}

	fmt.Printf("Task %d updated successfully.\n", ID)
	return nil
}

func (a *TaskServiceImpl) DeleteTask(ID int) error {
	// Call the repository to get the current task list
	data, err := a.Repository.ReadCurrentDatabase()
	if err != nil {
		return err
	}

	isFound := false
	// Find the task by ID and delete it
	for i := range data.Tasks {
		if data.Tasks[i].ID == ID {
			fmt.Printf("Delete task %d \"%s\"? (y/N) ", ID, data.Tasks[i].Title)

			var confirm string
			// Prompt the user for confirmation
			fmt.Scanln(&confirm)
			if confirm != "y" {
				fmt.Println("Task deletion cancelled.")
				return nil
			}

			// Delete the task from the slice
			data.Tasks = slices.Delete(data.Tasks, i, i+1)
			isFound = true
			break
		}
	}

	// Task not found validation
	if !isFound {
		fmt.Printf("Error: task %d not found.\n", ID)
		return nil
	}

	// Overwrite updated tasks back to file
	if err = a.Repository.OverwriteTask(data); err != nil {
		return err
	}

	fmt.Printf("Task %d deleted successfully.\n", ID)
	return nil
}
