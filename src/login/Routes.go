package login

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
}

func (r Route) String() string {
	return fmt.Sprintf("Route name:%s method:%s pattern:%s", r.Name, r.Method, r.Pattern)
}

type Routes []Route

var routes = Routes{
	Route{
		"Login",
		"GET",
		"/login",
		Login,
	},
	Route{
		"Reissue",
		"GET",
		"/login/reissue",
		Reissue,
	},
	Route{
		"Roles",
		"GET",
		"/login/roles",
		Roles,
	},
	Route{
		"APIStatus",
		"GET",
		"/login/status",
		util.StatsHandler,
	},
}
