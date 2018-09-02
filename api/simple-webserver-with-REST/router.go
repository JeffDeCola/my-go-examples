package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func jeffsRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		
		handler = route.RouteHandlerFunc

		// wrap in a logger before pasing to router
		handler = Logger(handler, route.RouteName)

		router.
			Methods(route.RouteMethod).
			Path(route.RoutePattern).
			Name(route.RouteName).
			Handler(handler)

	}
	return router
}
