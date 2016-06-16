package fox

import (
	"authz"
	"net/http"
	"util"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("FoxService")

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		authz.GetProvider().AddRestriction(route.Role, route.Method, route.Pattern)

		handler = route.HandlerFunc
		handler = util.NewTelemetry(handler, route.Name)

		handler = util.LoggingHandler(handler, log)
		handler = authz.PermissionHandler(handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		log.Debugf("Added route %s", route.String())
	}
	return router
}
