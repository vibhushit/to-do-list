package main

import (
	"fmt"
	"html/template"
	"net/http"
	"to-do-list/todo"
)

var tmpl *template.Template
var todoList *todo.ToDoList

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/tasks", tasksAPIHandler)
	http.HandleFunc("/api/add", addTaskAPIHandler)

	tmpl = template.Must(template.ParseFiles("../frontend/index.html"))

	todoList = todo.NewToDoList()

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

type ViewData struct {
	Tasks []todo.Task
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Tasks: todoList.ListTasks(),
	}
	tmpl.Execute(w, data)
}

func tasksAPIHandler(w http.ResponseWriter, r *http.Request) {
	tasks := todoList.ListTasks()
	respondJSON(w, tasks)
}

func addTaskAPIHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	taskTitle := r.FormValue("title")
	if taskTitle != "" {
		todoList.AddTask(taskTitle)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// You can use encoding/json to encode data and write it to the response writer.
}
