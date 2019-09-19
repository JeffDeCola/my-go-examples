// my-go-examples simple-blockchain.go

package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
)

// Block is your block
type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

// Message takes incoming JSON payload for writing data
type Message struct {
	Data string
}

var mutex = &sync.Mutex{}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func main() {

	fmt.Printf("\nSTARTING...\n")

	// Load port 8080 from your .env file
	err := godotenv.Load()
	checkErr(err)

	// Create a blockchain
	createBlockchain()

	// Start the web server
	go startWebServer()

	// Press return to exit
	fmt.Scanln()
	fmt.Println("EOF...")
}
