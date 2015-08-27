package main

import (
	"fmt"
	"time"
)

type Connection struct {
	Connection_ID string    `json:"connection_id"`
	Name          string    `json:"user_name"`
	Email         string    `json:"email"`
	Ages          int       `json:"ages"`
	Sex           string    `json:"sex"`
	Password      string    `json:"password"`
	Date_Created  time.Time `json:"date_created"`
}

func (c *Connection) valid() bool {
	return len(c.Connection_ID) > 0 && len(c.Name) > 0 && len(c.Email) > 0
}

func (c *Connection) printConnectionDetails() {

	fmt.Println("Name: ", c.Name)
	fmt.Println("Ages: ", c.Ages)

}

func (c *Connection) returnAgesIncremented(ages int) int {
	return (c.Ages + ages)
}

type Connections []Connection
