package main

import (
	"log"
	"net/http"
)

func main() {

	// Create router.
	myrouter := jeffsRouter()

	// Starts listening on localhost (127.0.0.1:1234)
	log.Fatal(http.ListenAndServe(":1234", myrouter))

}
