package config

import (
	"fmt"

	_ "github.com/lib/pq"
	mgo "gopkg.in/mgo.v2"
)

var DB *mgo.Database

var Todos *mgo.Collection

func init() {
	s, err := mgo.Dial("mongodb://tomoka:tomoka@localhost/todolist")
	if err != nil {
		panic(err)
	}
	if err = s.Ping(); err != nil {
		panic(err)
	}
	DB = s.DB("todolist")
	Todos = DB.C("todos")

	fmt.Println("you have successfully connected to your mongo db")
}
