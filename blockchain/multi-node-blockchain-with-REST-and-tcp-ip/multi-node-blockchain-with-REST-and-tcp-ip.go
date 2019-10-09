// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip.go

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

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func startWebServer() {

	// CREATE ROUTER
	myRouter := jeffsRouter()

	// LISTEN ON IP AND PORT
	fmt.Printf("\nListening on %s:%s\n\n", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, myRouter))

}

func main() {

	fmt.Printf("\nSTARTING...\n")

	// CREATE BLOCKCHAIN
	createBlockchain()

	// START WEBSERVER
	go startWebServer()

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Println("\n...DONE\n")
}
