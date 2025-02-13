package task

import (
	"errors"
	"fmt"
	"time"
)

type Status string

const (
	InProgress Status = "В процессе"
	Completed  Status = "Завершено"
)

type Task struct {
	ID        int
	Title     string
	Status    Status
	CreatedAt time.Time
}

type TaskManager struct {
	tasks []Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) AddTask(title string) Task {
	task := Task{
		ID:        len(tm.tasks) + 1,
		Title:     title,
		Status:    InProgress,
		CreatedAt: time.Now(),
	}
	tm.tasks = append(tm.tasks, task)
	return task
}

func (tm *TaskManager) GetTasks() []Task {
	return tm.tasks
}

func (tm *TaskManager) GetTaskByID(id int) (*Task, error) {
	for _, task := range tm.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (tm *TaskManager) UpdateTaskStatus(id int, status Status) error {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks[i].Status = status
			return nil
		}
	}
	return errors.New("task not found")
}

func (tm *TaskManager) DeleteTask(id int) error {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func (tm *TaskManager) PrintTasks() {
	if len(tm.tasks) == 0 {
		fmt.Println("Нет задач.")
		return
	}
	for _, task := range tm.tasks {
		fmt.Printf("ID: %d | Название: %s | Статус: %s | Создано: %s\n",
			task.ID, task.Title, task.Status, task.CreatedAt.Format(time.RFC1123))
	}
}
