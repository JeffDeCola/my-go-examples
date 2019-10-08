// tcp-ip server.go

package main

import (
	"fmt"
	"net"
)

func startTCPServer(ip string, port string) {

	// Listen for incoming connections
	fmt.Println("Listening on " + ip + ":" + port)
	listener, err := net.Listen("tcp", ip+":"+port)
	checkErr(err)
	defer listener.Close()

	// create a new connection each time we receive a connection request, and we need to serve it.
	// Server each connections concurrently
	for {
		// Wait for a connection request
		conn, err := listener.Accept()
		checkErr(err)

		go handleRequest(conn)
	}
}

// Handles incoming requests
func handleRequest(conn net.Conn) {

	// Make a buffer to hold incoming data
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer
	_, err := conn.Read(buf)
	checkErr(err)

	// Send a response back
	conn.Write([]byte("Message received."))
	// Close the connection
	conn.Close()
}
