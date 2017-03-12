package main

import (
	"io"
	"log"
	"net/http"
)

// response
func response(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {

	// Call the response function when you get a http request
	http.HandleFunc("/", response)

	// Starts listening on localhost (127.0.0.1:1234)
	log.Fatal(http.ListenAndServe(":1234", nil))

}
