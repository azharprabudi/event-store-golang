package storage

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// SimpleEventStoreStorageInterface ...
type SimpleEventStoreStorageInterface interface {
	FindAll(model []*interface{}, collectionName string, filter bson.M, sorts ...string) error
	FindOne(model *interface{}, collectionName string, filter bson.M, sorts ...string) error
	Create(model interface{}, collectionName string) error
	CountAll(collectionName string, filter bson.M) (int, error)
}

// SimpleEventStoreStorage ...
type SimpleEventStoreStorage struct {
	mongoDB *mgo.Database
}

// FindAll ...
func (esq *SimpleEventStoreStorage) FindAll(model []*interface{}, collectionName string, filter bson.M, sorts ...string) error {
	collSelected := esq.mongoDB.C(collectionName)
	err := collSelected.Find(filter).Sort(sorts...).All(&model)
	// check if there a query error, this function should be return error
	if err != nil {
		return err
	}
	return nil
}

// FindOne ...
func (esq *SimpleEventStoreStorage) FindOne(model *struct{}, collectionName string, filter bson.M, sorts ...string) error {
	collSelected := esq.mongoDB.C(collectionName)
	err := collSelected.Find(filter).Sort(sorts...).One(&model)
	// check if there a query error, this function should be return error
	if err != nil {
		return err
	}
	return nil
}

// Create ...
func (esq *SimpleEventStoreStorage) Create(model interface{}, collectionName string) error {
	collSelected := esq.mongoDB.C(collectionName)
	err := collSelected.Insert(model)
	// check if there a query error, this function should be return error
	if err != nil {
		return err
	}
	return nil
}

// CountAll ...
func (esq *SimpleEventStoreStorage) CountAll(collectionName string, filter bson.M) (int, error) {
	collSelected := esq.mongoDB.C(collectionName)
	count, err := collSelected.Find(filter).Count()
	// check if there a query error, this function should be return error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// NewSimpleEventStoreStorage ...
func NewSimpleEventStoreStorage(mongoDB *mgo.Database) SimpleEventStoreStorageInterface {
	return &SimpleEventStoreStorage{
		mongoDB: mongoDB,
	}
}
