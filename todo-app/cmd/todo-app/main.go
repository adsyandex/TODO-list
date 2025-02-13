package main

import (
	"fmt"
	"github.com/adsyandex/hello-world/todo-app/internal/menu"
	"github.com/adsyandex/hello-world/todo-app/internal/storage"
	"github.com/adsyandex/hello-world/todo-app/internal/task"
	"os"
)

func main() {
	fs := storage.NewFileStorage()
	tm := task.NewTaskManager()

	// Загрузка задач из файла
	tasks, err := fs.LoadTasks()
	if err == nil {
		for _, t := range tasks {
			tm.AddTask(t.Title) // Загрузка всех задач в менеджер
		}
	} else if !os.IsNotExist(err) {
		fmt.Println("Ошибка загрузки задач:", err)
	}

	// Запуск меню
	menu.ShowMenu(tm, fs)
}
