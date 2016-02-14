package fox

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func (r Route) String() string {
	return fmt.Sprintf("Route name:%s method:%s pattern:%s", r.Name, r.Method, r.Pattern)
}

type Routes []Route

var routes = Routes{
	Route{
		"GetFox",
		"GET",
		"/fox/foxes/{foxId}",
		Show,
	},
	Route{
		"GetFoxes",
		"GET",
		"/fox/foxes",
		List,
	},
	Route{
		"UpdateFox",
		"PUT",
		"/fox/foxes/{foxId}",
		Update,
	},
	Route{
		"AddFox",
		"POST",
		"/fox/foxes",
		Add,
	},
	Route{
		"DeleteFox",
		"DELETE",
		"/fox/foxes/{foxId}/delete",
		Delete,
	},
	Route{
		"APIStatus",
		"GET",
		"/fox/status",
		Stats,
	},
}
