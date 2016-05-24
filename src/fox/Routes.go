package fox

import (
	"fmt"
	"net/http"
	"util"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Role		string
}

func (r Route) String() string {
	return fmt.Sprintf("Route name:%s method:%s pattern:%s role:%s", r.Name, r.Method, r.Pattern, r.Role)
}

type Routes []Route

var routes = Routes{
	Route{
		"GetFox",
		"GET",
		"/fox/foxes/{foxId}",
		Show,
		"registry user",
	},
	Route{
		"GetFoxes",
		"GET",
		"/fox/foxes",
		List,
		"registry user",
	},
	Route{
		"UpdateFox",
		"PUT",
		"/fox/foxes/{foxId}",
		Update,
		"registry administrator",
	},
	Route{
		"AddFox",
		"POST",
		"/fox/foxes",
		Add,
		"registry administrator",
	},
	Route{
		"DeleteFox",
		"DELETE",
		"/fox/foxes/{foxId}",
		Delete,
		"registry administrator",
	},
	Route{
		"APIStatus",
		"GET",
		"/fox/status",
		util.StatsHandler,
		"registry user",
	},
}
