package repository

import "github.com/mhaatha/go-learn/mini-project/tasker/internal/model"

type TaskRepository interface {
	ReadCurrentDatabase() (model.JSONFormat, error)
	OverwriteTask(currentDatabase model.JSONFormat) error
}
