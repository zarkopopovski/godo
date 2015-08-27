package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

func OpenConnectionSession() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func (c *Connection) CreateNewConnection() bool {
	session := OpenConnectionSession()
	defer session.Close()

	collection := session.DB("godo").C("userdata")
	err := collection.Insert(c)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (t *Todo) CreateNewToDO() {
	session := OpenConnectionSession()
	defer session.Close()

	c := session.DB("godo").C("usertasks")
	err := c.Insert(t)
	if err != nil {
		log.Fatal(err)
	}
}

func (t *Todo) DeleteTask() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	//c := session.DB("godo").C("usertasks")

}
