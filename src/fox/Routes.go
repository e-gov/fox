package main

import (
	"net/http"
    "fmt"
)

type Route struct {
	Name			string
	Method			string
	Pattern 		string
	HandlerFunc 	http.HandlerFunc
}

func (r Route) String() string{
    return fmt.Sprintf("Route name:%s method:%s pattern:%s", r.Name, r.Method, r.Pattern)
}

type Routes []Route


var routes = Routes{
	Route{
		"GetFox",
		"GET",
		"/fox/foxes/{foxId}",
		FoxShow,
	},
    Route{
		"GetFoxes",
		"GET",
		"/fox/foxes",
	FoxList,   
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
	Route{
		"DeleteFox",
		"GET",
		"/fox/foxes/{foxId}/delete",
		DeleteFox,
	},
	Route{
		"APIStatus",
		"GET",
		"/fox/status",
		ShowStats,
	},
    
}

