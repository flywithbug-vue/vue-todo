package model

import (
	"time"
	"todo-go/core/mongo"

	"gopkg.in/mgo.v2/bson"

	"github.com/flywithbug/log4go"
)

const (
	IncrementTodo  = "todo"
	todoCollection = "todo"
)

type Todo struct {
	Id        int64  `json:"id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	Destroy   bool   `json:"destroy,omitempty" bson:"destroy,omitempty"`
}

func InsertTodo(t *Todo) error {
	incrementId, err := mongo.GetIncrementId(IncrementTodo)
	if err != nil {
		log4go.Error(err.Error())
		return err
	}
	t.Completed = false
	t.CreatedAt = time.Now().Unix()
	t.UpdatedAt = t.CreatedAt
	t.Id = incrementId
	return mongo.Insert(db, todoCollection, t)
}

func FindAllTodos() ([]Todo, error) {
	var results []Todo
	err := mongo.FindAll(db, todoCollection, bson.M{"destroy": false}, nil, &results)
	return results, err
}

func FindTodoById(id int64) (*Todo, error) {
	todo := new(Todo)
	err := mongo.FindOne(db, todoCollection, bson.M{"_id": id}, nil, todo)
	return todo, err
}

func (t *Todo) Remove(id int64) error {
	return mongo.Remove(db, todoCollection, bson.M{"_id": id})
}
