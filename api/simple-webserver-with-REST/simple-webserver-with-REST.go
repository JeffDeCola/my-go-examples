package main

import (
	"fmt"
	"log"
	"net/http"
)

func startServer() {

	// Create a router
	myRouter := jeffsRouter()

	// Start listening on localhost (127.0.0.1:1234)
	log.Fatal(http.ListenAndServe(":1234", myRouter))
}

func main() {

	// Starts listening on localhost (127.0.0.1:1234)
	go startServer()

	// Press retrun to exit
	fmt.Scanln()
	fmt.Println("done")

}
