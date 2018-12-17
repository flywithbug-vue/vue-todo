package model

var db = "todpo"

func SetDBName(dbName string) {
	db = dbName
}

func DBName() string {
	return db
}

type ModelOperation interface {
	FindAll() ([]interface{}, error)
	Insert() error
	Update(id string) error
	Remove(id string) error
	FindById(id string) (interface{}, error)
}
