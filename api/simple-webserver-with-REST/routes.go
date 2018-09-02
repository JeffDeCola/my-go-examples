package main

import "net/http"

// Route - The struct for the routes/endpoints
type Route struct {
	RouteName        string
	RouteMethod      string
	RoutePattern     string
	RouteHandlerFunc http.HandlerFunc
}

// Routes is slice
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodosIndex",
		"GET",
		"/todos",
		TodosIndex,
	},
	Route{
		"TodosCreate",
		"POST",
		"/todos",
		TodosCreate,
	},
	Route{
		"TodosShow",
		"GET",
		"/todos/{todoID}",
		TodosShow,
	},
	Route{
		"GetPeopleEndpoint",
		"GET",
		"/people",
		GetPeopleEndpoint,
	},
	Route{
		"GetPersonEndpoint",
		"GET",
		"/people/{id}",
		GetPersonEndpoint,
	},
	Route{
		"CreatePersonEndpoint",
		"POST",
		"/people/{id}",
		CreatePersonEndpoint,
	},
	Route{
		"DeletePersonEndpoint",
		"DELETE",
		"/people/{id}",
		DeletePersonEndpoint,
	},
}
