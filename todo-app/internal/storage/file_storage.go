package storage

import (
	"encoding/json"
	"github.com/adsyandex/hello-world/todo-app/internal/task"
	"os"
)

const storageFile = "tasks.json"

type FileStorage struct {
}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (fs *FileStorage) SaveTasks(tasks []task.Task) error {
	file, err := os.Create(storageFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func (fs *FileStorage) LoadTasks() ([]task.Task, error) {
	file, err := os.Open(storageFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []task.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
