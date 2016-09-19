package fox

import (
	"net/http"

	"github.com/e-gov/fox/authz"
	"github.com/e-gov/fox/util"

	"github.com/gorilla/mux"
	log "github.com/Sirupsen/logrus"
)



func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		var ap = authz.GetProvider()
		if ap != nil {
			ap.AddRestriction(route.Role, route.Method, route.Pattern)
		}

		handler = route.HandlerFunc
		handler = util.NewTelemetry(handler, route.Name)

		handler = authz.PermissionHandler(handler)
		handler = util.LoggingHandler(handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		log.WithFields(log.Fields{
			"path": route.Pattern,
			"method": route.Method,
		}).Infof("Added route %s", route.String())
	}
	return router
}
