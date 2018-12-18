package model

import (
	"errors"
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
	UpdatedAt int64  `json:"updated_at,omitempty" bson:"updated_at"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at"`
	Destroy   bool   `json:"destroy,omitempty" bson:"destroy"`
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

//selector bson.M{"updated_at": 0, "created_at": 0} value值为0时不获取值
func FindAllTodos() ([]Todo, error) {
	var results []Todo
	err := mongo.FindAll(db, todoCollection, bson.M{"destroy": false}, nil, &results)
	return results, err
}

// query :bson.M{"destroy": false}「filter」. selector：bson.M{"updated_at": 0, "created_at": 0}value值为0时不获取值
func FindTodos(query, selector interface{}) ([]Todo, error) {
	var results []Todo
	err := mongo.FindAll(db, todoCollection, query, selector, &results)
	return results, err
}

func FindTodoById(id int64) (*Todo, error) {
	todo := new(Todo)
	err := mongo.FindOne(db, todoCollection, bson.M{"_id": id}, nil, todo)
	return todo, err
}

func (t *Todo) Remove() error {
	if t.Id == 0 {
		return errors.New("item id can not be 0")
	}
	t.UpdatedAt = time.Now().Unix()
	return mongo.Remove(db, todoCollection, bson.M{"_id": t.Id})
}

func (t *Todo) DestroyT() error {
	if t.Id == 0 {
		return errors.New("item id can not be 0")
	}
	t.UpdatedAt = time.Now().Unix()
	return mongo.Update(db, todoCollection, bson.M{"_id": t.Id}, bson.M{"$set": bson.M{"destroy": true, "updated_at": t.UpdatedAt}})
}

func (t *Todo) Update() error {
	if t.Id == 0 {
		return errors.New("item id can not be 0")
	}
	t.UpdatedAt = time.Now().Unix()
	return mongo.Update(db, todoCollection, bson.M{"_id": t.Id}, t)
}

func CheckAllTodoItems(complete bool) error {
	updatedAt := time.Now().Unix()
	selector := bson.M{"completed": !complete, "destroy": false}
	data := bson.M{"$set": bson.M{"completed": complete, "updated_at": updatedAt}}
	_, err := mongo.UpdateAll(db, todoCollection, selector, data)
	return err
}
