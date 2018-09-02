package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const htmlCoreyIndex = `
<html>
<body>
<a href="/jeff">Goto /jeff</a>
</body>
</html>
`

// rootHandler
func rootHandler(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	io.WriteString(res, "hello, world!\n")
}

// jeffHandler
func jeffHandler(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	io.WriteString(res, "hello, Jeff!\n")
}

// coreyHandler
func coreyHandler(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	fmt.Fprintf(res, htmlCoreyIndex)
}

func main() {

	// Call your function when you get a http request for "/" or "/jeff"
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/jeff", jeffHandler)
	http.HandleFunc("/corey", coreyHandler)

	// Starts listening on localhost (127.0.0.1:1234)
	log.Fatal(http.ListenAndServe(":1234", nil))

}
