// my-go-examples simple-tcp-ip-server.go

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
)

const (
	ip   = "localhost"
	port = "3333"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
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

func open(addr string) (*bufio.ReadWriter, error) {

	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "Dialing "+addr+" failed")
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}

func sendString(data string) {

	rw, err := open(ip + port)
	checkErr(err)

	log.Println("Send the string request.")
	_, err = rw.WriteString("STRING\n")
	checkErr(err)

	log.Println("Send Additional data.")
	_, err = rw.WriteString("Additional data.\n")
	checkErr(err)

	log.Println("Flush the buffer.")
	err = rw.Flush()
	checkErr(err)
}

func main() {

	go startTCPServer(serverIP, serverPort)

	go startTCPClient(remoteIP, remotePort)

}
