package storage

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// EventStoreQueryInterface ...
type EventStoreQueryInterface interface {
	FindAll(model []*interface{}, collectionName string, filter bson.M, sorts ...string) error
	FindOne(model *interface{}, collectionName string, filter bson.M, sorts ...string) error
	Create(model interface{}, collectionName string) error
	CountAll(collectionName string, filter bson.M) (int, error)
}

// EventStoreQuery ...
type EventStoreQuery struct {
	mongoDB *mgo.Database
}

// FindAll ...
func (esq *EventStoreQuery) FindAll(model []*interface{}, collectionName string, filter bson.M, sorts ...string) error {
	collSelected := esq.mongoDB.C(collectionName)
	err := collSelected.Find(filter).Sort(sorts...).All(&model)
	// check if there a query error, this function should be return error
	if err != nil {
		return err
	}
	return nil
}

// FindOne ...
func (esq *EventStoreQuery) FindOne(model *interface{}, collectionName string, filter bson.M, sorts ...string) error {
	collSelected := esq.mongoDB.C(collectionName)
	err := collSelected.Find(filter).Sort(sorts...).One(&model)
	// check if there a query error, this function should be return error
	if err != nil {
		return err
	}
	return nil
}

// Create ...
func (esq *EventStoreQuery) Create(model interface{}, collectionName string) error {
	collSelected := esq.mongoDB.C(collectionName)
	err := collSelected.Insert(model)
	// check if there a query error, this function should be return error
	if err != nil {
		return err
	}
	return nil
}

// CountAll ...
func (esq *EventStoreQuery) CountAll(collectionName string, filter bson.M) (int, error) {
	collSelected := esq.mongoDB.C(collectionName)
	count, err := collSelected.Find(filter).Count()
	// check if there a query error, this function should be return error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// NewEventStoreQuery ...
func NewEventStoreQuery(mongoDB *mgo.Database) EventStoreQueryInterface {
	return &EventStoreQuery{
		mongoDB: mongoDB,
	}
}
