package main

import (
	"net/http"
)

type Route struct {
	Name			string
	Method			string
	Pattern 		string
	HandlerFunc 	http.HandlerFunc
}

type Routes []Route


var routes = Routes{
	Route{
		"FoxList",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"GetFox",
		"GET",
		"/fox/foxes/{foxId}",
		FoxShow,
	},
	Route{
		"AddFox",
		"POST",
		"/fox/add",
		AddFox,
	},
	Route{
		"UpdateFox",
		"POST",
		"/fox/foxes/{foxId}",
		UpdateFox,
	},
}