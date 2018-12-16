package model

import (
	"errors"
	"time"
	"todo-go/core/mongo"
)

const (
	STATUS_LOGIN    = 1
	STATUS_LOGOUT   = 2
	loginCollection = "login"
)

type Login struct {
	//Id 			bson.ObjectId
	UserId     string `bson:"user_id"`     // 用户ID
	Token      string `bson:"token"`       // 用户TOKEN
	CreateTime int64  `bson:"create_time"` // 登录日期
	LoginIp    string `bson:"login_ip"`    // 登录IP
	Status     int    `bson:"status"`      //status 1 已登录，0表示退出登录
	Forbidden  bool   `bson:"forbidden"`   //false 表示未禁言
	userAgent  string `bson:"user_agent"`  //用户UA
}

func UserLogin(userId, userAgent string) (l *Login, err error) {
	l = new(Login)
	l.UserId = userId
	l.userAgent = userAgent
	//l.Token = genToken()
	l.CreateTime = time.Now().Unix()
	l.Status = 1
	err = l.Insert()
	return
}

func (l Login) FindAll() ([]Login, error) {
	var results []Login
	err := mongo.FindAll(db, userCollection, nil, nil, &results)
	return results, err
}

func (l Login) Insert() error {
	if l.UserId == "" {
		return errors.New("user_id can not be nil")
	}
	return mongo.Insert(db, loginCollection, l)
}

func (l Login) Update(id string) error {
	panic("implement me")
}

func (l Login) Remove(id string) error {
	panic("implement me")
}

func (l Login) FindById(id string) (interface{}, error) {
	panic("implement me")
}
