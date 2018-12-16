package mongo

import (
	"github.com/flywithbug/log4go"
	"gopkg.in/mgo.v2"
)

var globalS *mgo.Session

func DialMgo(url string)  {
	s, err := mgo.Dial(url)
	if err != nil {
		log4go.Fatal("create session error ", err)
	}
	globalS = s
	log4go.Info("mongodb connected")
}

func connect(db,collection string)(*mgo.Session,*mgo.Collection)  {
	if globalS == nil {
		log4go.Fatal("mgo disconnected")
	}
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s,c
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func IsExist(db, collection string, query interface{}) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, results interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(results)
}

func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(selector,update)
}

func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(selector)
}
