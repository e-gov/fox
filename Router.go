package main

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
 	"github.com/gorilla/handlers"
)

func NewRouter() *mux.Router{
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
	}
	return router
}
