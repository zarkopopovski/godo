package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	ages := r.FormValue("ages")
	sex := r.FormValue("sex")
	password := r.FormValue("password")

	userAges, _ := strconv.Atoi(ages)

	connection := &Connection{Name: name, Email: email, Ages: userAges, Sex: sex, Password: password, Date_Created: time.Now()}
	CreateNewConnection(connection)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	connection := LoginWithCredentials(email, password)
	userID := fmt.Sprintf("%x", string(connection.Id))
	log.Println(userID)

	if connection != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		userMap := make(map[string]string)
		userMap["token"] = userID

		if err := json.NewEncoder(w).Encode(userMap); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func AddTask(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	todoTitle := r.FormValue("todo_title")
	todoBody := r.FormValue("todo_body")
	//todoDueDate := r.FormValue("due_date")

	if len(token) == 0 || len(todoTitle) == 0 {
		return
	}

	todo := &Todo{UserID: bson.ObjectIdHex(token), Title: todoTitle, Body: todoBody, Completed: false, Due: time.Now()}
	CreateNewTask(todo)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	todoID := r.FormValue("todo_id")

	DeleteTask(token, todoID)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	todoID := r.FormValue("todo_id")
	todoTitle := r.FormValue("todo_title")
	todoBody := r.FormValue("todo_body")

	todoCompleted, _ := strconv.ParseBool(r.FormValue("complete"))
	//todoDueDate := r.FormValue("due_date")

	todo := &Todo{
		Id:        bson.ObjectIdHex(todoID),
		UserID:    bson.ObjectIdHex(token),
		Title:     todoTitle,
		Body:      todoBody,
		Completed: todoCompleted,
		Due:       time.Now()}

	UpdateExistingTask(todo)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	todos := ListAllTasks(token)

	if len(todos) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todos); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
