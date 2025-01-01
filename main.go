package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var data TodoPageData

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println(r.FormValue("title"))
		title := r.FormValue("title")
		if title != "" {
			data.Todos = append(data.Todos, Todo{Title: title, Done: false})
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func markTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println(r.FormValue("title"))
		title := r.FormValue("title")
		if title != "" {
			for i, todo := range data.Todos {
				if todo.Title == title {
					data.Todos[i].Done = true
				}
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println(r.FormValue("title"))
		title := r.FormValue("title")
		if title != "" {
			for i, todo := range data.Todos {
				if todo.Title == title {
					data.Todos = append(data.Todos[:i], data.Todos[i+1:]...)
				}
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func main() {
	data = TodoPageData{
		PageTitle: "Day TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/add", addTaskHandler)
	http.HandleFunc("/mark", markTaskHandler)
	http.HandleFunc("/delete", deleteTaskHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
