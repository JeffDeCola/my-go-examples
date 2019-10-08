package main

import "net/http"

// Route - The struct for the route endpoints (e.g. /jeff)
type Route struct {
	RouteName        string
	RouteHTTPVerb    string
	RouteEndPoint    string
	RouteHandlerFunc http.HandlerFunc
}

// Routes is slice
type Routes []Route

var routes = Routes{
	Route{
		"GetIndex",
		"GET",
		"/",
		rootHandler,
	},
	Route{
		"GetData",
		"GET",
		"/getdata/{todoID}",
		getdataHandler,
	},
	Route{
		"PostData",
		"POST",
		"/postdata/{todoID}",
		postdataHandler,
    },
    Route{
		"PutData",
		"PUT",
		"/putdata/{todoID}",
		putdataHandler,
    },
    Route{
		"DeleteData",
		"DELETE",
		"/deletedata/{todoID}",
		deletedataHandler,
	},
}
