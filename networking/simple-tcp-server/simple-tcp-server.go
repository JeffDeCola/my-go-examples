// my-go-examples simple-tcp-server.go

package main

import (
	"fmt"
	"log"
	"net"
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

func startTCPServer() {

	// LISTEN ON IP AND PORT
	fmt.Printf("\nListening on %s:%s\n\n", ip, port)
	server, err := net.Listen("tcp", ip+":"+port)
	checkErr(err)
	defer server.Close()

	// CREATE A CONNECTION FOR EACH REQUEST
	// Serve connections concurrently
	for {

		// Wait for a connection request
		conn, err := server.Accept()
		checkErr(err)

		log.Println("Opening a connection")
		log.Println("----------------------------------------------------------------")

		go handleRequest(conn)
	}
}

func main() {

	// START TCP SERVER
	go startTCPServer()

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Println("done")

}
