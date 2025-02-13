package storage

import "github.com/adsyandex/hello-world/todo-app/internal/task"

type Storage interface {
	SaveTask(task.Task) error
	GetTasks() ([]task.Task, error)
	GetTaskByID(id int) (task.Task, error)
	UpdateTask(task.Task) error
	DeleteTask(id int) error
	GetNextID() int
}
