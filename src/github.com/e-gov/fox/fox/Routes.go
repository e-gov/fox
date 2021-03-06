package fox

import (
	"fmt"
	"net/http"

	"github.com/e-gov/fox/util"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Role        string
}

func (r Route) String() string {
	return fmt.Sprintf("Route name:%s method:%s pattern:%s role:%s", r.Name, r.Method, r.Pattern, r.Role)
}

type Routes []Route

// The last element of the route is internal role name that is required for access
// * indicates an un-restricted route.
var routes = Routes{
	Route{
		"GetFox",
		"GET",
		"/fox/foxes/{foxId}",
		Show,
		"*",
	},
	Route{
		"GetFoxes",
		"GET",
		"/fox/foxes",
		List,
		"*",
	},
	Route{
		"UpdateFox",
		"PUT",
		"/fox/foxes/{foxId}",
		Update,
		"ADMIN",
	},
	Route{
		"AddFox",
		"POST",
		"/fox/foxes",
		Add,
		"ADMIN",
	},
	Route{
		"DeleteFox",
		"DELETE",
		"/fox/foxes/{foxId}",
		Delete,
		"ADMIN",
	},
	Route{
		"APIStatus",
		"GET",
		"/fox/status",
		util.StatsHandler,
		"*",
	},
}
