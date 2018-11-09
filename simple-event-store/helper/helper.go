package helper

import (
	"errors"

	"github.com/globalsign/mgo"
)

// OpenDBConnection ...
func OpenDBConnection(dbSource string, dbName string) (*mgo.Database, error) {
	// validate dbSource
	if dbSource == "" {
		return nil, errors.New("db source cannot be empty")
	}

	// validate dbName
	if dbName == "" {
		return nil, errors.New("no database selected")
	}

	// connect to dbsource
	dbConn, err := mgo.Dial(dbSource)
	if err != nil {
		return nil, err
	}

	return dbConn.DB(dbName), nil
}
