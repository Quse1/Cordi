/*1. Добавление задач (Create):
Пользователь вводит название новой задачи.
Задача добавляется в список.
Приложение поддерживает динамическое добавление произвольного количества задач.

2. Чтение списка задач (Read):
Отображает список всех текущих задач с указанием их номеров.

3. Обновление задач (Update):
Пользователь вводит имя задачи, которую нужно изменить.
Если задача найдена, приложение предлагает ввести новое имя.
Новое имя должно быть длиной не менее 3 символов.
Обновленная задача сохраняется в списке.

4. Удаление задач (Delete):
Пользователь вводит имя задачи, которую необходимо удалить.
Если задача найдена, она удаляется из списка.
Приложение сообщает о успешном удалении.

5. Обработка ошибок:
Если введена несуществующая команда, выводится сообщение о неверной команде.
При попытке обновить или удалить несуществующую задачу пользователю предлагается повторить ввод.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	NewTask := make([]string, 0)
	var inxUpdate int = -1
	var inxDelete int = -1

	for {
		fmt.Println("Enter your command (create, read, update, delete): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "create" {
			fmt.Println("Enter task name:")
			reader := bufio.NewReader(os.Stdin)
			Task, _ := reader.ReadString('\n')
			Task = strings.TrimSpace(Task)
			NewTask = append(NewTask, Task)
		} else if input == "read" {
			for i, t := range NewTask {
				fmt.Printf("%d. %s\n", i+1, t)
			}
		} else if input == "update" {
			fmt.Println("Enter task name to update:")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			for i, t := range NewTask {
				if input == t {
					inxUpdate = i
					fmt.Println("Enter new task name:")
					break
				}
			}
			if inxUpdate == -1 {
				fmt.Println("Invalid name. Please try again.")
				continue
			}
			reader = bufio.NewReader(os.Stdin)
			updTask, _ := reader.ReadString('\n')
			updTask = strings.TrimSpace(updTask)

			if len(updTask) < 3 {
				fmt.Println("The new task name is too short! Please, try again.")
				continue
			}

			NewTask[inxUpdate] = updTask
			fmt.Printf("Updated task #%v with name \"%s\" successfully!\n", inxUpdate, updTask)

		} else if input == "delete" {
			fmt.Println("Enter task name to remove:")
			reader := bufio.NewReader(os.Stdin)
			delTask, _ := reader.ReadString('\n')
			delTask = strings.TrimSpace(delTask)
			for i, t := range NewTask {
				if delTask == t {
					inxDelete = i
					fmt.Println("Enter task name:")
					break
				}
			}
			if inxDelete == -1 {
				fmt.Println("Invalid name. Please try again.")
				continue
			}
			oldTask := NewTask[inxDelete]
			NewTask = append(NewTask[:inxDelete], NewTask[inxDelete+1:]...)
			fmt.Printf("Removed task #%v with name \"%s\" successfully!\n", inxDelete, oldTask)
		} else {
			fmt.Println("Invalid command! Please, try again!")
		}
	}

}
