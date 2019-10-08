// my-go-examples simple-blockchain router.go

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)
// Message takes incoming JSON payload for writing data
type Message struct {
	Data string  `json:"data"`
}

func jeffsRouter() *mux.Router {

	// MAKE ROUTER
	router := mux.NewRouter().StrictSlash(true)

	// LOAD ROUTES ONE BY ONE
	for _, route := range routes {

		var handler http.Handler
		handler = route.RouteHandlerFunc

		// WRAP IN LOGGER
		handler = Logger(handler, route.RouteName)

		router.
			Name(route.RouteName).
			Methods(route.RouteHTTPVerb).
			Path(route.RouteEndPoint).
			Handler(handler)
	}

	return router
}
