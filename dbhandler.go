package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func CreateNewConnection(c *Connection) bool {
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

func CreateNewToDO(t *Todo) {
	session := OpenConnectionSession()
	defer session.Close()

	c := session.DB("godo").C("usertasks")
	err := c.Insert(t)
	if err != nil {
		log.Fatal(err)
	}
}

func ListAllTasks(token string) []Todo {
	session := OpenConnectionSession()
	defer session.Close()

	var todos []Todo

	c := session.DB("godo").C("usertasks")
	err := c.Find(bson.M{"user_id": bson.ObjectIdHex(token)}).All(&todos)
	if err != nil {
		log.Fatal(err)
	}

	return todos
}

func LoginWithCredentials(email string, password string) *Connection {
	session := OpenConnectionSession()
	defer session.Close()

	var connection *Connection

	c := session.DB("godo").C("userdata")
	err := c.Find(
		bson.M{"email": email,
			"$and": []interface{}{
				bson.M{"password": password},
			}}).One(&connection)
	if err != nil {
		log.Fatal(err)
	}

	return connection
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
