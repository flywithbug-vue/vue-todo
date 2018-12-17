package mongo

import (
	"github.com/flywithbug/log4go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionsIncrementIds = "increment_ids"
)

type incrementIds struct {
	Id            string `json:"_id" bson:"_id"`                       //自增标记Id
	SequenceValue int64  `json:"sequence_value" bson:"sequence_value"` //自增 id值
}

func CreateIncrementIds(db, incrementName string) error {
	cou := incrementIds{
		Id:            incrementName,
		SequenceValue: 1,
	}
	return Insert(db, collectionsIncrementIds, cou)
}

func GetIncrementIdByName(db, incrementName string) (int64, error) {
	increment := new(incrementIds)
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"sequence_value": 1}},
		ReturnNew: true,
	}
	ms, c := connect(db, collectionsIncrementIds)
	defer ms.Close()
	_, err := c.Find(bson.M{"_id": incrementName}).Apply(change, increment)
	if err != nil {
		err = CreateIncrementIds(db, incrementName)
		if err != nil {
			log4go.Error(err.Error())
			return -1, err
		}
		increment.SequenceValue = 1
	}
	return increment.SequenceValue, nil
}
