package task

import (
 "testing"
 "time"
)

func TestAddTask(t *testing.T) {
 tm := NewTaskManager()
 task := tm.AddTask("Тестовая задача")

 if task.ID != 1 {
 

 t.Errorf("Ожидался ID задачи 1, но получили %d", task.ID)
 }
 if task.Title != "Тестовая задача" {
  t.Errorf("Ожидался заголовок 'Тестовая задача', но получили '%s'", task.Title)
 }
}

func TestUpdateTaskStatus(t *testing.T) {
 tm := NewTaskManager()
 task := tm.AddTask("Тестовая задача")

 if err := tm.UpdateTaskStatus(task.ID, Completed); err != nil {
  t.Errorf("Ошибка обновления статуса: %v", err)
 }

 if tm.tasks[0].Status != Completed {
  t.Errorf("Ожидался статус 'Завершено', но получили '%s'", tm.tasks[0].Status)
 }
}
