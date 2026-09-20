package service

type TaskService interface {
	AddTask(title string) error
	ListTasks() error
	MarkTaskAsDone(ID int) error
	ShowTask(ID int) error
	EditTask(ID int, title string) error
	DeleteTask(ID int) error
}
