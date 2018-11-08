package eventstore

import (
	"event-store/SimpleEventStore/services"

	"github.com/globalsign/mgo"
)

// SimpleEventStore ...
type SimpleEventStore struct {
	service services.EventStoreServiceInterface
}

// NewSimpleventStore ...
func NewSimpleventStore(maxLimitAction int, collectionName string, mongoDB *mgo.Database) *SimpleEventStore {
	return &SimpleEventStore{
		service: services.NewEventStoreService(maxLimitAction, collectionName, mongoDB),
	}
}
