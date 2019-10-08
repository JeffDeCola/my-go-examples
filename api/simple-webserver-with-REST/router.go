package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

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
