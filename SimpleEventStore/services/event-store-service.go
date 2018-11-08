package services

import (
	"event-store/SimpleEventStore/model"
	"event-store/SimpleEventStore/storage"
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// EventStoreServiceInterface ...
type EventStoreServiceInterface interface {
	Get() (*bson.M, error)
	Insert() error
}

// EventStoreService ...
type EventStoreService struct {
	maxLimitAction int
	collections    eventStoreCollections
	storage        storage.EventStoreQueryInterface
}

// eventStoreCollections ...
type eventStoreCollections struct {
	events        string
	snapshot      string
	configuration string
}

// Get ...
func (es *EventStoreService) Get() (*bson.M, error) {
	/* should be return data like model parsed */
	return nil, nil
}

// Insert ...
func (es *EventStoreService) Insert(model interface{}) error {
	// if the configuration doesnt exists add to the configuration
	var err error
	err = es.initConfiguration()
	if err != nil {
		return err
	}

}

/*

this function should be used for initial configuration if the collection configuration doesnt exits.

*/
func (es *EventStoreService) initConfiguration() error {
	// get count data to check if collection is exists
	count, err := es.storage.CountAll(es.collections.configuration, nil)
	if err != nil {
		return err
	}

	// if current count is zero, than added the data to the collection configuration
	if count == 0 {
		es.storage.Create(model.ConfPropertyCollection{
			DefaultMaxLimitAction: es.maxLimitAction,
			CurrentMaxLimitAction: es.maxLimitAction,
		})
	}
	return nil
}

// NewEventStoreService ...
func NewEventStoreService(maxLimitAction int, collectionName string, mongoDB *mgo.Database) EventStoreServiceInterface {
	eventCollecName := fmt.Sprintf("%s_events", collectionName)
	snapshotCollecName := fmt.Sprintf("%s_snapshots", collectionName)
	configurationCollecName := fmt.Sprintf("%s_configuration", collectionName)

	return &EventStoreService{
		storage:        storage.NewEventStoreQuery(mongoDB),
		maxLimitAction: maxLimitAction,
		collections: eventStoreCollections{
			events:        eventCollecName,
			snapshot:      snapshotCollecName,
			configuration: configurationCollecName,
		},
	}
}
