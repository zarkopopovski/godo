# GoDO

## GoDO with GO

**GoDO** is basic "ToDo" web service implementation in Go. Its developed on top on JSON Restful Api example from [The New Stack](http://thenewstack.io/make-a-restful-json-api-go/) article. Each action is implemented as different api call using POST method, response is returned as JSON data.

Available functions:

* SignIn
* SignOut
* Login
* AddTask
* RemoveTask
* UpdateTask
* ListTasks

**GoDO** use MongoDB NoSQL database as storage for users and tasks. Users and tasks will be stored in Users collections where each task will be stored in User object per user for easy quering. Future optimizations will be made.

## Building
**GoDO** is tested on Mac OSx 10.10.5

	go get
	go build

## Running
**GoDO** can be runned from source or if is deployed as single executable like console application.
	
	go run *.go
	or
	./GoDO

## Project Status
This project will be actively developed as opensource software.