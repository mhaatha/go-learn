package repository

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"io"
	"os"

	"github.com/mhaatha/go-learn/mini-project/tasker/internal/model"
)

func NewTaskRepository(fileName string) TaskRepository {
	return &TaskRepositoryImpl{
		fileName: fileName,
	}
}

type TaskRepositoryImpl struct {
	fileName string
}

func (a *TaskRepositoryImpl) ReadCurrentDatabase() (model.JSONFormat, error) {
	// Open the database file
	file, err := os.OpenFile(a.fileName, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return model.JSONFormat{}, err
	}
	defer file.Close()

	// Decode the raw JSON data and unmarshal it into the data struct
	dec := jsontext.NewDecoder(file)
	var data model.JSONFormat
	for {
		if err := json.UnmarshalDecode(dec, &data); err != nil {
			if err == io.EOF {
				break
			} else {
				return model.JSONFormat{}, err
			}
		}
	}

	return data, nil
}

func (a *TaskRepositoryImpl) OverwriteTask(currentDatabase model.JSONFormat) error {
	// Marshal the updated data back into JSON
	jsonData, err := json.Marshal(currentDatabase, jsontext.Multiline(true))
	if err != nil {
		return err
	}

	// Overwrite the database file with the updated JSON data
	err = os.WriteFile(a.fileName, jsonData, 0600)
	if err != nil {
		return err
	}

	return nil
}
