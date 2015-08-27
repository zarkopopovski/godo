package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Connection struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name         string        `json:"user_name"`
	Email        string        `json:"email"`
	Ages         int           `json:"ages"`
	Sex          string        `json:"sex"`
	Password     string        `json:"password"`
	Date_Created time.Time     `json:"date_created"`
}

func (c *Connection) valid() bool {
	return len(c.Id) > 0 && len(c.Name) > 0 && len(c.Email) > 0
}

func (c *Connection) printConnectionDetails() {

	fmt.Println("Name: ", c.Name)
	fmt.Println("Ages: ", c.Ages)

}

func (c *Connection) returnAgesIncremented(ages int) int {
	return (c.Ages + ages)
}

type Connections []Connection
