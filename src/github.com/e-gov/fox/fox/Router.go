package fox

import (
	"net/http"

	"github.com/e-gov/fox/authz"
	"github.com/e-gov/fox/util"

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

		handler = authz.PermissionHandler(handler)
		handler = util.LoggingHandler(handler, log)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		log.Debugf("Added route %s", route.String())
	}
	return router
}
