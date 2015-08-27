package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"SignIn", "POST", "/signin", SignIn},
	Route{"SignOut", "POST", "/signout", SignOut},
	Route{"Login", "POST", "/login", Login},
	Route{"AddTask", "POST", "/addtask", AddTask},
	Route{"RemoveTask", "POST", "/removetask", RemoveTask},
	Route{"UpdateTask", "POST", "/updatetask", UpdateTask},
	Route{"ListTasks", "POST", "/listtasks", ListTasks},
}
