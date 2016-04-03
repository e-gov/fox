package fox

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"authz"
	"github.com/rcrowley/go-metrics"
	"time"
	thatlog "log"
	"os"
)

var log = logging.MustGetLogger("FoxService")

func NewRouter(name string) *mux.Router{
	nodeName = name
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		
		authz.GetProvider().AddRestriction(route.Role, route.Method, route.Pattern)
			
		handler = route.HandlerFunc
		handler = NewTelemetry(handler, route.Name)
		
		handler = LoggingHandler(handler, log)
		handler = PermissionHandler(handler)
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
		log.Debugf("Added route %s", route.String())
		
	}
	go metrics.Log(metrics.DefaultRegistry, 30 * time.Second, thatlog.New(os.Stderr, "metrics: ", thatlog.Lmicroseconds))
	return router
}
