package model

import (
	"encoding/json"
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
	tJson, _ := json.Marshal(todos)
	fmt.Println(string(tJson))

	//todo, err = FindTodoById(2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(todo)
}

func TestUserFunctions(t *testing.T) {
	mongo.DialMgo("127.0.0.1:27017")

	user := new(User)
	user.RealName = "ori"
	user.Password = "pass"
	user.Account = "ori"
	user.Title = "CEO"
	user.Phone = "129"
	user.Mail = "admin@admin.com"
	user.Sex = 1

	err := user.Insert()
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}
