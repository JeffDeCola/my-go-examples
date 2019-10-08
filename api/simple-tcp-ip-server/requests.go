// my-go-examples simple-tcp-ip-server requests.go

package main

import (
	"bufio"
	"io"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

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

		// CALL HANDLER
		// ADD OR SUBSTRACT
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
