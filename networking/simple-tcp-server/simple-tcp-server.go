// my-go-examples simple-tcp-server.go

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
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

// handleInt
func handleAdd(rw *bufio.ReadWriter) {

	s := "Please enter two integers"
	returnMessage(s, rw)

	// WAITING FOR INT 1
	s1, err := rw.ReadString('\n')
	checkErr(err)
	s1 = strings.Trim(s1, "\n ")
	i1, _ := strconv.Atoi(s1)
	s = "Received first integer: " + s1
	returnMessage(s, rw)

	// WAITING FOR INT 2
	s2, err := rw.ReadString('\n')
	checkErr(err)
	s2 = strings.Trim(s2, "\n ")
	i2, _ := strconv.Atoi(s2)
	s = "Received second integer: " + s2
	returnMessage(s, rw)

	// THE SUM
	sum := i1 + i2
	s = "The sum of " + s1 + " + " + s2 + " = " + strconv.Itoa(sum)
	returnMessage(s, rw)

}

func handleSubstract(rw *bufio.ReadWriter) {

	s := "Please enter two integers"
	returnMessage(s, rw)

	// WAITING FOR INT 1
	s1, err := rw.ReadString('\n')
	checkErr(err)
	s1 = strings.Trim(s1, "\n ")
	i1, _ := strconv.Atoi(s1)
	s = "Received first integer: " + s1
	returnMessage(s, rw)

	// WAITING FOR INT 2
	s2, err := rw.ReadString('\n')
	checkErr(err)
	s2 = strings.Trim(s2, "\n ")
	i2, _ := strconv.Atoi(s2)
	s = "Received second integer: " + s2
	returnMessage(s, rw)

	// SUBTRACT
	sum := i1 - i2
	s = "Subtracting " + s1 + " - " + s2 + " = " + strconv.Itoa(sum)
	returnMessage(s, rw)

}

func startServer() {

	// LISTEN ON IP AND PORT
	fmt.Printf("\nListening on %s:%s\n\n", ip, port)
	server, err := net.Listen("tcp", ip+":"+port)
	checkErr(err)
	defer server.Close()

	// CREATE CONNECTION FOR EACH REQUEST
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

func handleRequest(conn net.Conn) {

	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	// READ FROM CONN UTIL EOF
	for {

		s := "Waiting for command: ADD, SUBTRACT or EOF"
		returnMessage(s, rw)

		cmd, err := rw.ReadString('\n')
		// TRIM CMD
		cmd = strings.Trim(cmd, "\n ")

		s = "Received command: " + cmd
		returnMessage(s, rw)

		// CHECK FOR EOF
		switch {
		case err == io.EOF:
			s = "Reached EOF"
			returnMessage(s, rw)
			s = "Closing this connection"
			returnMessage(s, rw)
			log.Println("----------------------------------------------------------------")
			return
		case err != nil:
			log.Println("ERROR reading command. Got: ", cmd, err)
			return
		}

		// CALL HANDLER ADD OR SUBSTRACT
		// Otherwise close connection
		switch {
		case cmd == "ADD":
			handleAdd(rw)
		case cmd == "SUBTRACT":
			handleSubstract(rw)
		case cmd == "EOF":
			s = "Received EOF"
			returnMessage(s, rw)
			s = "Closing this connection"
			returnMessage(s, rw)
			log.Println("----------------------------------------------------------------")
			return
		default:
			s = "Did not get correct command. Received: " + cmd
			returnMessage(s, rw)
			s = "Closing this connection"
			returnMessage(s, rw)
			log.Println("----------------------------------------------------------------")
			return
		}
	}

}
func returnMessage(s string, rw *bufio.ReadWriter) {
	log.Println(s)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
}

func main() {

	// START WEBSERVER
	go startServer()

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Println("done")

}
