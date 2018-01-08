package todos

import (
	"errors"
	"net/http"

	"github.com/Tomoka64/todoWithMongoDB/config"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Isbn  string
	Title string
	Due   string
}

func AllTodos() ([]Todo, error) {
	tds := []Todo{}
	err := config.Todos.Find(bson.M{}).All(&tds)
	if err != nil {
		return nil, err
	}
	return tds, nil
}

func OneTodo(r *http.Request) (Todo, error) {
	td := Todo{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return td, errors.New("400. Bad Request.")
	}
	err := config.Todos.Find(bson.M{"isbn": isbn}).One(&td)
	if err != nil {
		return td, err
	}
	return td, nil
}

func PutTodo(r *http.Request) (Todo, error) {
	td := Todo{}
	td.Isbn = r.FormValue("isbn")
	td.Title = r.FormValue("title")
	td.Due = r.FormValue("due")
	if td.Isbn == "" || td.Title == "" || td.Due == "" {
		return td, errors.New("All fields must be complete.")
	}
	err := config.Todos.Insert(td)
	if err != nil {
		return td, errors.New("Internal Error")
	}
	return td, nil
}

func UpdateTodo(r *http.Request) (Todo, error) {
	td := Todo{}
	td.Isbn = r.FormValue("isbn")
	td.Title = r.FormValue("title")
	td.Due = r.FormValue("due")
	if td.Isbn == "" || td.Title == "" || td.Due == "" {
		return td, errors.New("All fields must be complete.")
	}
	err := config.Todos.Update(bson.M{"isbn": td.Isbn}, &td)
	if err != nil {
		return td, err
	}
	return td, nil
}

func DeleteTodo(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("Bad REquest")
	}
	err := config.Todos.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("INTernal S Eror")
	}
	return nil
}
