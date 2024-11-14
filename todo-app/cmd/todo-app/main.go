package main

import (
 "fmt"
 "os"
 "todo-app/internal/storage"
 "todo-app/internal/task"
 "todo-app/internal/menu"
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
