package main

import (
	//"encoding/json"
	"fmt"

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
	connection.CreateNewConnection()
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
}

func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	todoTitle := r.FormValue("todo_title")
	todoBody := r.FormValue("todo_body")
	//todoDueDate := r.FormValue("due_date")

	if len(token) == 0 || len(todoTitle) == 0 {
		return
	}

	todo := &Todo{Title: todoTitle, Body: todoBody, Completed: false, Due: time.Now()}
	todo.CreateNewToDO()
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	//token := r.FormValue("token")
	//todoID := r.FormValue("todo_id")

	//DeleteTask(token, todoID)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
}
