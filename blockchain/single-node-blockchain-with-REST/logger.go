// my-go-examples single-node-blockchain-with-REST logger.go

package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Logger passed handler and name, get more info
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
