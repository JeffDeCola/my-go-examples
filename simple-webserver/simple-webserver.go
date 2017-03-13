package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// rootHandler
func rootHandler(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	io.WriteString(res, "hello, world!\n")
	fmt.Printf("res is: %+v\n\n", res)
}

//jeff Handler
func jeffHandler(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	io.WriteString(res, "hello, Jeff!\n")
	fmt.Printf("res is: %+v\n\n", res)
}

func main() {

	// Call the response function when you get a http request for "/"
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/jeff", jeffHandler)

	// Starts listening on localhost (127.0.0.1:1234)
	log.Fatal(http.ListenAndServe(":1234", nil))

}
