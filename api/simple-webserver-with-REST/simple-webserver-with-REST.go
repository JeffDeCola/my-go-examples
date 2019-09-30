package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	ip   = "127.0.0.1"
	port = "1234"
)

func startServer() {

	// Create a router
	myRouter := jeffsRouter()

	// LISTEN ON IP AND PORT
	fmt.Printf("\nListening on %s:%s\n\n", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, myRouter))

}

func main() {

	go startServer()

	// Press retrun to exit
	fmt.Scanln()
	fmt.Println("done")

}
