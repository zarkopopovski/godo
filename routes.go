package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func RoutesMap(api *ApiConnection) Routes {
	var routes = Routes{
		Route{"Index", "GET", "/", api.Index},
		Route{"SignIn", "POST", "/signin", api.SignIn},
		Route{"SignOut", "POST", "/signout", api.SignOut},
		Route{"Login", "POST", "/login", api.Login},
		Route{"AddTask", "POST", "/addtask", api.AddTask},
		Route{"RemoveTask", "POST", "/removetask", api.RemoveTask},
		Route{"UpdateTask", "POST", "/updatetask", api.UpdateTask},
		Route{"ListTasks", "POST", "/listtasks", api.ListTasks},
	}

	return routes
}
