package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type MongoConnection struct {
	dbSession *mgo.Session
}

func OpenConnectionSession() (mongoConnection *MongoConnection) {
	mongoConnection = new(MongoConnection)
	mongoConnection.createNewDBConnection()

	return
}

func (mConnection *MongoConnection) createNewDBConnection() (err error) {
	mConnection.dbSession, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	mConnection.dbSession.SetMode(mgo.Monotonic, true)

	return
}

func (mConnection *MongoConnection) CreateNewConnection(c *Connection) bool {

	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	collection := session.DB("godo").C("userdata")
	err := collection.Insert(c)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (mConnection *MongoConnection) CreateNewTask(t *Todo) bool {

	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	c := session.DB("godo").C("usertasks")
	err := c.Insert(t)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (mConnection *MongoConnection) ListAllTasks(token string) []Todo {

	if mConnection.dbSession == nil {
		return nil
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	var todos []Todo

	c := session.DB("godo").C("usertasks")
	err := c.Find(bson.M{"user_id": bson.ObjectIdHex(token)}).All(&todos)
	if err != nil {
		log.Fatal(err)
	}

	return todos
}

func (mConnection *MongoConnection) LoginWithCredentials(email string, password string) *Connection {

	if mConnection.dbSession == nil {
		return nil
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	var connection *Connection

	c := session.DB("godo").C("userdata")
	err := c.Find(bson.M{"email": email, "$and": []interface{}{bson.M{"password": password}}}).One(&connection)
	if err != nil {
		log.Fatal(err)
	}

	return connection
}

func (mConnection *MongoConnection) DeleteTask(token string, todoID string) bool {

	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	c := session.DB("godo").C("usertasks")
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(todoID), "$and": []interface{}{bson.M{"user_id": bson.ObjectIdHex(token)}}})

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (mConnection *MongoConnection) UpdateExistingTask(todo *Todo) bool {

	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	c := session.DB("godo").C("usertasks")
	err := c.UpdateId(todo.Id, todo)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
