package model

import (
	"errors"
	"time"
	"todo-go/core/mongo"

	"gopkg.in/mgo.v2/bson"
)

const (
	STATUS_LOGIN    = 1
	STATUS_LOGOUT   = 2
	loginCollection = "login"
)

type Login struct {
	UserId     string `bson:"user_id"`     // 用户ID
	Token      string `bson:"token"`       // 用户TOKEN
	CreateTime int64  `bson:"create_time"` // 登录日期
	LoginIp    string `bson:"login_ip"`    // 登录IP
	Status     int    `bson:"status"`      //status 1 已登录，0表示退出登录
	Forbidden  bool   `bson:"forbidden"`   //false 表示未禁言
	userAgent  string `bson:"user_agent"`  //用户UA
	UpdatedAt  int64  `json:"updated_at,omitempty" bson:"updated_at"`
}

func UserLogin(userId, userAgent, token, ip string) (l *Login, err error) {
	l = new(Login)
	l.UserId = userId
	l.userAgent = userAgent
	l.Token = token
	l.CreateTime = time.Now().Unix()
	l.UpdatedAt = l.CreateTime
	l.Status = 1
	l.LoginIp = ip
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

//status 0 退出登录，1 登录
//	return mongo.Update(db, todoCollection, bson.M{"_id": t.Id}, bson.M{"$set": bson.M{"title": t.Title, "completed": t.Completed, "updated_at": t.UpdatedAt}})
func UpdateLoginStatus(token string, status int) error {
	updateAt := time.Now().Unix()
	return mongo.Update(db, loginCollection, bson.M{"token": token}, bson.M{"$set": bson.M{"status": status, "updated_at": updateAt}})
}

func FindLoginByToken(token string) (l *Login, err error) {
	l = new(Login)
	err = mongo.FindOne(db, loginCollection, bson.M{"token": token}, nil, &l)
	return
}
