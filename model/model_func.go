package model

const (
	db = "todo-go"
)

type ModelOperation interface {
	FindAll() ([]interface{}, error)
	Insert() error
	Update(id string) error
	Remove(id string) error
	FindById(id string) (interface{}, error)
}
