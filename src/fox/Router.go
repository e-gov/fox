package fox

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("FoxService")

func NewRouter(name string) *mux.Router{
	nodeName = name
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = StatHandler(handler, route.Name)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handlers.LoggingHandler(os.Stdout, handler))
		log.Debugf("Added route %s", route.String())
	}
	return router
}
