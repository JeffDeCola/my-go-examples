package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// rootHandler
func rootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("%+v\n", req)
	io.WriteString(res, "hello, world!\n")
}

// jeffHandler
func jeffHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("%+v\n", req)
	io.WriteString(res, "hello, Jeff!\n")
}

func main() {

	// Call the response function when you get a http request
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/jeff", jeffHandler)

	// Starts listening on localhost (127.0.0.1:1234)
	log.Fatal(http.ListenAndServe(":1234", nil))

}
