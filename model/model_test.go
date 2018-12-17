package model

import (
	"fmt"
	"testing"
	"todo-go/core/mongo"
)

func TestTodo(t *testing.T) {
	mongo.DialMgo("127.0.0.1:27017")

	todo := new(Todo)
	todo.Title = "test"
	todo.Completed = false
	err := InsertTodo(todo)
	if err != nil {
		panic(err)
	}
	todos, err := FindAllTodos()
	if err != nil {
		panic(err)
	}

	fmt.Println(todos)

	//todo, err = FindTodoById(2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(todo)
}
