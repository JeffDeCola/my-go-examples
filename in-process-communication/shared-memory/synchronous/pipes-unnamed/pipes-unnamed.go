package main

import (
	"fmt"
	"io"
)

// sendDataLoop will write to the pipe 3 times
func sendDataLoop(pw *io.PipeWriter, data string) {

	n := 1
	for n <= 3 {
		data := fmt.Sprintf("%s %d", data, n)
		_, err := pw.Write([]byte(data))
		if err != nil {
			fmt.Printf("error with pr.Write method: %v \n", err)
		}
		fmt.Printf("SEND       : %s \n", data)
		n++
	}

	pw.Close()
	fmt.Println("CLOSED PIPE")

}

// rcvData will read from the pipe 50 bytes at a time until EOF
func rcvDataUntilEOF(pr *io.PipeReader) {

	for {
		rcvData := make([]byte, 20)
		_, err := pr.Read(rcvData)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Printf("error with pr.Read method: %v \n", err)
		}
		fmt.Printf("RECEIVED   : %s \n", rcvData)
	}

}

func main() {

	// CREATE UNNAMED PIPE
	pr, pw := io.Pipe()

	// DATA TO SEND
	data := "I am the data that will be sent"

	// SEND - WRITE TO SHARED MEMORY
	go sendDataLoop(pw, data)

	// RECEIVE - READ FROM SHARED MEMORY
	rcvDataUntilEOF(pr)

	fmt.Println("DONE")

}
