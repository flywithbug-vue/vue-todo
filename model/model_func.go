package model

import "todo-go/core/mongo"

var db = "todo"

func SetDBName(dbName string) {
	db = dbName
	mongo.SetIncrementDBName(dbName)
}

func init() {
	mongo.SetIncrementDBName(db)
}

func DBName() string {
	return db
}

type OperationModel interface {
	FindAll() ([]interface{}, error)
	Insert() error
	Update(id string) error
	Remove(id string) error
	FindById(id string) (interface{}, error)
}
