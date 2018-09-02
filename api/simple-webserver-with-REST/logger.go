package main

import (
	"log"
	"net/http"
	"time"
)

// Logger passed handler and name
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()

		inner.ServeHTTP(res, req)

		log.Printf(
			"%s\t%s\t%s\t%s",
			req.Method,
			req.RequestURI,
			name,
			time.Since(start),
		)
	})
}
