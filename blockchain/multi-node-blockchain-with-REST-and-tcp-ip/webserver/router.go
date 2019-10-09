// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip router.go

package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Message takes incoming JSON payload for writing data
type Message struct {
	Data string  `json:"data"`
}

// JeffsRouter is the router
func JeffsRouter() *mux.Router {

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
