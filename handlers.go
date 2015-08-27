package main

import (
	//"encoding/json"
	"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	//"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	fmt.Println(r.FormValue("email"))
	fmt.Fprint(w, "Welcome!\n"+r.FormValue("email"))
	//var err jsonErr
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
	//var err jsonErr
}

func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
	//var err jsonErr
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
	//var err jsonErr
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
	//var err jsonErr
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
	//var err jsonErr
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["token"])
	//var err jsonErr
}
