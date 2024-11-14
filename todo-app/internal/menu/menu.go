package menu

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "todo-app/internal/task"
 "todo-app/internal/storage"
)

func ShowMenu(tm *task.TaskManager, fs *storage.FileStorage) {
 reader := bufio.NewReader(os.Stdin)

 for {
  fmt.Println("\nМеню:")
  fmt.Println("1. Просмотр задач")
  fmt.Println("2. Добавить задачу")
  fmt.Println("3. Обновить статус задачи")
  fmt.Println("4. Удалить задачу")
  fmt.Println("5. Выйти")
  fmt.Print("Выберите опцию: ")

  option, _ := reader.ReadString('\n')
  option = option[:len(option)-1]

  switch option {
  case "1":
   tm.PrintTasks()
  case "2":
   fmt.Print("Введите название задачи: ")
   title, _ := reader.ReadString('\n')
   title = title[:len(title)-1]
   task := tm.AddTask(title)
   fmt.Printf("Задача добавлена: %s\n", task.Title)
  case "3":
   fmt.Print("Введите ID задачи для обновления: ")
   idStr, _ := reader.ReadString('\n')
   id, err := strconv.Atoi(idStr[:len(idStr)-1])
   if err != nil {
    fmt.Println("Неверный ID.")
    continue
   }
   fmt.Print("Введите новый статус (1 - В процессе, 2 - Завершено): ")
   statusStr, _ := reader.ReadString('\n')
   status, _ := strconv.Atoi(statusStr[:len(statusStr)-1])
   var statusEnum task.Status
   if status == 1 {
    statusEnum = task.InProgress
   } else if status == 2 {
    statusEnum = task.Completed
   } else {
    fmt.Println("Неверный статус.")
    continue
   }
   err = tm.UpdateTaskStatus(id, statusEnum)
   if err != nil {
    fmt.Println("Ошибка:", err)
   } else {
    fmt.Println("Статус задачи обновлен.")
   }
  case "4":
   fmt.Print("Введите ID задачи для удаления: ")
   idStr, _ := reader.ReadString('\n')
   id, err := strconv.Atoi(idStr[:len(idStr)-1])
   if err != nil {
    fmt.Println("Неверный ID.")
    continue
   }
   err = tm.DeleteTask(id)
   if err != nil {
    fmt.Println("Ошибка:", err)
   } else {
    fmt.Println("Задача удалена.")
   }
  case "5":
   tasks := tm.GetTasks()
   fs.SaveTasks(tasks)
   fmt.Println("До свидания!")
   return
  default:
   fmt.Println("Неверный выбор. Попробуйте снова.")
  }
 }
}
