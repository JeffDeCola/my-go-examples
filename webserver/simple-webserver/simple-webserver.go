// my-go-examples simple-webserver.go

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	ip   = "127.0.0.1"
	port = "1234"
)

const htmlMonkeyIndex = `
<html>
<body>
<p> You are in the monkeyHandler </p>
<a href="/jeff">Goto /jeff</a>
</body>
</html>
`

// rootHandler
func rootHandler(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	io.WriteString(res, "You are in the rootHandler\n")
	io.WriteString(res, "hello, world!\n")
}

// jeffHandler
func jeffHandler(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	io.WriteString(res, "You are in the jeffHandler\n")
	io.WriteString(res, "hello, Jeff!\n")
}

// monkeyHandler
func monkeyHandler(res http.ResponseWriter, req *http.Request) {
	// print out the req struct for fun
	fmt.Printf("req is %+v\n\n", req.URL)
	fmt.Fprintf(res, htmlMonkeyIndex)
}

func startServer() {
	log.Fatal(http.ListenAndServe(ip+":"+port, nil))
}

func main() {

	// Call your function when you get a http request for "/" or "/jeff"
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/jeff", jeffHandler)
	http.HandleFunc("/monkey", monkeyHandler)

	// Starts listening on localhost (127.0.0.1:1234)
	go startServer()

	// Press retrun to exit
	fmt.Scanln()
	fmt.Println("done")

}
