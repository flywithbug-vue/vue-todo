package model

import (
	"errors"
	"strings"
	"todo-go/core/mongo"

	"gopkg.in/mgo.v2/bson"
)

const (
	userCollection = "user"
)

type User struct {
	Id       int64  `json:"_id" bson:"_id"`
	UserId   string `json:"user_id" bson:"user_id"`
	Account  string `json:"account"`
	Password string
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"` // 0保密，1男 2女
	RealName string `json:"real_name" bson:"real_name"`
	Title    string `json:"title"`
}

func FindAllUsers() ([]User, error) {
	var results []User
	err := mongo.FindAll(db, userCollection, nil, nil, &results)
	return results, err
}

func (u *User) Insert() error {
	if mongo.IsExist(db, userCollection, bson.M{"account": u.Account}) {
		return errors.New("account 已存在")
	}
	if mongo.IsExist(db, userCollection, bson.M{"mail": u.Mail}) {
		return errors.New("mail 已存在")
	}
	if u.UserId == "" {
		u.Id, _ = mongo.GetIncrementId("user")
		u.UserId = bson.NewObjectId().Hex()
	}
	return mongo.Insert(db, userCollection, u)
}

func (u *User) CheckLogin(account, pass string) (err error) {
	err = mongo.FindOne(db, userCollection, bson.M{"account": account, "password": pass}, nil, u)
	return err
}

func (u *User) Update(id string) error {
	if !strings.EqualFold(u.UserId, id) {
		return errors.New("userId not equal id")
	}
	return mongo.Update(db, userCollection, bson.M{"_id": bson.ObjectIdHex(id)}, u)
}

func (u User) Remove(id string) error {
	return mongo.Remove(db, userCollection, bson.M{"_id": bson.ObjectIdHex(id)})
}

func (u *User) FindById(id string) error {
	err := mongo.FindOne(db, userCollection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &u)
	return err
}

func AddAdminUser() error {
	u := new(User)
	u.Account = "admin"
	u.Password = "flywithbug123"
	u.Mail = "flywithbug@gmail.com"
	u.Title = "admin"
	u.Phone = "phone"
	u.RealName = "Jack"
	u.Sex = 1
	return u.Insert()
}
