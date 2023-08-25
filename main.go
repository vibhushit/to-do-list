package main

import (
	"fmt"
	"to-do-list/todo"
)

func main() {
	fmt.Println("Welcome to the to-do list app...")

	//code
	todoList := todo.NewToDoList()

	todoList.AddTask("Buy Groceries")
	todoList.AddTask("Write code")
	todoList.CompleteTask(1)

	fmt.Println(todoList.ListTasks())

	fmt.Println("Goodbye!")
}
