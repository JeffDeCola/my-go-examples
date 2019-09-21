// my-go-examples simple-blockchain.go

package main

import (
	"fmt"
	"log"
)

const (
	httpPort = "8080"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func main() {

	fmt.Printf("\nSTARTING...\n")

	// Create a blockchain
	createBlockchain()

	// Start the web server
	go startWebServer()

	// Press return to exit
	fmt.Scanln()
	fmt.Println("DONE")
}
