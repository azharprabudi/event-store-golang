package main

import (
	"time"

	ses "github.com/event-store/simple-event-store/"
	sesHelper "github.com/event-store/simple-event-store/helper"
)

// ExampleUserBalance ...
type ExampleUserBalance struct {
	UserID    string     `bson:"userId"`
	Balance   int64      `bson:"balance"`
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt"`
}

func main() {
	helper.OpenDBConnection
	db, err := sesHelper.OpenDBConnection("mongodb://root:root123456@ds115022.mlab.com:15022/cqrs", "cqrs")
	if err != nil {
		panic(err)
	}

	newSes := ses.NewSimpleEventStore(10, "user_balance", db)

	exampleUserBalance := ExampleUserBalance{
		Balance:   64000,
		UserID:    "8937ca65-2cf9-41fc-b23f-21be09e55938",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: nil,
	}

}
