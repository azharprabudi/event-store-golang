package main

import (
	ses "event-store/simple-event-store"
	sesHelper "event-store/simple-event-store/helper"
	"time"
)

// ExampleUserBalance ...
type ExampleUserBalance struct {
	UserID    string     `bson:"userId"`
	Balance   int64      `bson:"balance"`
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt"`
}

func main() {
	db, err := sesHelper.OpenDBConnection("mongodb://root:root123456@ds115022.mlab.com:15022/cqrs", "cqrs")
	if err != nil {
		panic(err)
	}
	newSes := ses.NewSimpleEventStore(10, "user_balance", db, []string{"userId", "balance", "createdAt", "updatedAt"})
	newSes.Insert(ExampleUserBalance{
		UserID:    "8937ca65-2cf9-41fc-b23f-21be09e55938",
		Balance:   64000,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: nil,
	})
}
