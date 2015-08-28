package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Todo struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserID    bson.ObjectId `bson:"user_id" json:"user_id"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	Completed bool          `json:"completed"`
	Due       time.Time     `json:"due"`
}

func (t *Todo) valudate() bool {
	return len(t.Id) > 0 && len(t.Title) > 0
}

type Todos []Todo
