package mongo

import (
	"gopkg.in/mgo.v2"
)

var globalS *mgo.Session

func DialMgo(url string) {
	s, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	if globalS == nil {
		panic("mgo disconnected")
	}
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
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
	return c.Update(selector, update)
}

/*
selector := bson.M{"name": "Tom"}
data := bson.M{"$set": bson.M{"age": 22}}
*/
func UpdateAll(db, collection string, selector, data interface{}) (*mgo.ChangeInfo, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	changInfo, err := c.UpdateAll(selector, data)
	return changInfo, err
}

func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	_, err := c.RemoveAll(selector)
	return err
}
