package main

import (
	"net/http"

	"github.com/Tomoka64/todoWithMongoDB/todos"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/todos", todos.Index)
	http.HandleFunc("/todos/show", todos.Show)
	http.HandleFunc("/todos/create", todos.Create)
	http.HandleFunc("/todos/create/process", todos.CreateProcess)
	http.HandleFunc("/todos/update", todos.Update)
	http.HandleFunc("/todos/update/process", todos.UpdateProcess)
	http.HandleFunc("/todos/delete/process", todos.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}
