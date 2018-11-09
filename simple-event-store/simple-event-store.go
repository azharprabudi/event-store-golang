package service

import (
	"event-store/simple-event-store/storage"
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// SimpleEventStorserviceInterface ...
type SimpleEventStorserviceInterface interface {
	Get() (*bson.M, error)
	Insert(model interface{}) error
}

// SimpleEventStorseservice ...
type SimpleEventStorseservice struct {
	maxLimitAction      int
	collections         SimpleEventStoreCollections
	storage             storage.SimpleEventStorsestorageInterface
	propertyCollections []string
}

// SimpleEventStoreCollections ...
type SimpleEventStoreCollections struct {
	events        string
	snapshot      string
	configuration string
}

// Get ...
func (ses *SimpleEventStorseservice) Get() (*bson.M, error) {
	/* should be return data like model parsed */
	return nil, nil
}

// Insert ...
func (ses *SimpleEventStorseservice) Insert(model interface{}) error {
	// if the configuration dosesnt exists add to the configuration
	fmt.Println(model)
	// var err error
	// err = ses.initConfiguration()
	// if err != nil {
	// 	return err
	// }
	// return err
	return nil
}

/*

this function should be used for initial configuration if the collection configuration dosesnt exits.

*/
func (ses *SimpleEventStorseservice) initConfiguration() error {
	// get count data to check if collection is exists
	count, err := ses.storage.CountAll(ses.collections.configuration, nil)
	if err != nil {
		return err
	}

	// if current count is zero, than added the data to the collection configuration
	if count == 0 {
		// ses.storage.Create(model.ConfPropertyCollection{
		// 	DefaultMaxLimitAction: ses.maxLimitAction,
		// 	CurrentMaxLimitAction: ses.maxLimitAction,
		// })
	}
	return nil
}

/*
get bson file from property collections
*/
func (ses *)

/*

NewSimpleEventStore simple explain:
@maxLimitAction => for snapshoting when the action reach the limit
@collectionName => specify this event store focus on spesific collection
@mongoDB => thats a connector mongodb
@props => list of property want to save in the mongodb collection

*/
func NewSimpleEventStore(maxLimitAction int, collectionName string, mongoDB *mgo.Database, propertyCollections []string) SimpleEventStorserviceInterface {
	eventCollecName := fmt.Sprintf("%s_events", collectionName)
	snapshotCollecName := fmt.Sprintf("%s_snapshots", collectionName)
	configurationCollecName := fmt.Sprintf("%s_configuration", collectionName)

	return &SimpleEventStorseservice{
		storage:        storage.NewSimpleEventStorsestorage(mongoDB),
		maxLimitAction: maxLimitAction,
		collections: SimpleEventStoreCollections{
			events:        eventCollecName,
			snapshot:      snapshotCollecName,
			configuration: configurationCollecName,
		},
		propertyCollections: propertyCollections,
	}
}
