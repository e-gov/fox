package login

import (
	"net/http"
	"os"

	"github.com/e-gov/fox/util"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("LoginService")

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = util.NewTelemetry(handler, route.Name)

		handler = util.LoggingHandler(handler, log)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handlers.LoggingHandler(os.Stdout, handler))
		log.Debug("Added " + route.String())

	}
	return router
}
