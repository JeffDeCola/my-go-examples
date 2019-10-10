// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip requests.go

package tcpserver

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// HandleRequest handles TCP requests
func HandleRequest(conn net.Conn) {

	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	s := "Opening a connection"
	returnMessage(s, rw)
	s = "----------------------------------------------------------------"
	returnMessage(s, rw)

	// READ FROM CONN UTIL EOF
	for {

		s := "Waiting for command: ADDNEWBLOCK or EOF"
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
			s = "----------------------------------------------------------------"
			returnMessage(s, rw)
			return
		case err != nil:
			s = "ERROR reading command. Got: " + cmd
			returnMessage(s, rw)
			return
		}

		// CALL HANDLER
		// ADDNEWBLOCK
		// Otherwise close connection
		switch {
		case cmd == "ADDNEWBLOCK":
			handleAddNewBlock(rw)
		// case cmd == "NEWCOMMAND":
		//	handleSubstract(rw)
		case cmd == "EOF":
			s = "Received EOF"
			returnMessage(s, rw)
			s = "Closing this connection"
			returnMessage(s, rw)
			s = "----------------------------------------------------------------"
			returnMessage(s, rw)
			return
		default:
			s = "Did not get correct command. Received: " + cmd
			returnMessage(s, rw)
			s = "Closing this connection"
			returnMessage(s, rw)
			s = "----------------------------------------------------------------"
			returnMessage(s, rw)
			return
		}
	}
}

func returnMessage(s string, rw *bufio.ReadWriter) {
	log.Println("TCPSERVER:      " + s)
	_, err := rw.WriteString("--- " + s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
}
