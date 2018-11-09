package service

import (
	"fmt"

	sesModel "github.com/event-store/simple-event-store/model"
	sesStorage "github.com/event-store/simple-event-store/storage"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// BaseModel ...
type BaseModel struct {
	ID    bson.ObjectId `bson:"_id"`
	Model interface{}   `bson:"model"`
}

// SimpleEventStorserviceInterface ...
type SimpleEventStorserviceInterface interface {
	Get() (*bson.M, error)
	Insert(model interface{}) error
}

// SimpleEventStoreservice ...
type SimpleEventStoreservice struct {
	maxLimitAction      int
	collections         SimpleEventStoreCollections
	storage             sesStorage.SimpleEventStoreStorageInterface
	propertyCollections []string
}

// SimpleEventStoreCollections ...
type SimpleEventStoreCollections struct {
	events        string
	snapshot      string
	configuration string
}

// Get ...
func (ses *SimpleEventStoreservice) Get() (*bson.M, error) {
	/* should be return data like model parsed */
	return nil, nil
}

// Insert ...
func (ses *SimpleEventStoreservice) Insert(model interface{}) error {
	// logger struct will be added
	// result := initConfiguration() correctly using in initialization, to create insert a new table, and create collectionz
	loggerID := bson.NewObjectId()
	mLogger := BaseModel{
		ID:    loggerID,
		Model: model,
	}
	// added to events
	err := ses.storage.Create(mLogger, ses.collections.events)
	if err != nil {
		return err
	}

	// check if the configuration collection reach the maxmimum limit
	currCountLimitation := new(sesModel.ConfPropertyCollection)
	ses.storage.FindOne(currCountLimitation, ses.collections.configuration, nil)
	if currCountLimitation.CurrentMaxLimitAction-1 == 0 {

	}

	// fmt.Println(model)
	// // var err error
	// // err = ses.initConfiguration()
	// // if err != nil {
	// // 	return err
	// // }
	// // return err
	// return nil
}

/*

this function should be used for initial configuration if the collection configuration dosesnt exits.

*/
func (ses *SimpleEventStoreservice) initConfiguration() error {
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

NewSimpleEventStore simple explain:
@maxLimitAction => for snapshoting when the action reach the limit
@collectionName => specify this event store focus on spesific collection
@mongoDB => thats a connector mongodb

*/
func NewSimpleEventStore(maxLimitAction int, collectionName string, mongoDB *mgo.Database, propertyCollections []string) SimpleEventStorserviceInterface {
	eventCollecName := fmt.Sprintf("%s_events", collectionName)
	snapshotCollecName := fmt.Sprintf("%s_snapshots", collectionName)
	configurationCollecName := fmt.Sprintf("%s_configuration", collectionName)

	return &SimpleEventStoreservice{
		storage:        sesStorage.NewSimpleEventStoreStorage(mongoDB),
		maxLimitAction: maxLimitAction,
		collections: SimpleEventStoreCollections{
			events:        eventCollecName,
			snapshot:      snapshotCollecName,
			configuration: configurationCollecName,
		},
	}
}
