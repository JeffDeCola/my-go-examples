package main

import (
	"fmt"
	"io"
	"time"
)

// sendDataLoop will write to the pipe
func sendDataLoop(pw *io.PipeWriter, data string) {

	n := 1
	for n <= 5 {
		data := fmt.Sprintf("%s %d", data, n)
		_, err := pw.Write([]byte(data))
		if err != nil {
			fmt.Printf("error with pr.Write method: %v", err)
		}
		fmt.Println("SENT DATA   : " + data)
		n++
	}

	pw.Close()
	fmt.Println("CLOSED PIPE")

}

// rcvData will read from the pipe 50 bytes at a time
func rcvData(pr *io.PipeReader) (string, error) {

	buffer := make([]byte, 20)

	_, err := pr.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("error with pr.Read method: %v", err)
	}

	return string(buffer), err

}

func main() {

	// CREATE UNNAMED PIPE
	pr, pw := io.Pipe()

	// DATA TO SEND
	data := "I am the data that will be sent"
	fmt.Printf("Data to Send: %s \n", data)

	// SEND
	// WILL WRITE TO SHARED MEMORY IN A LOOP AND BLOCK UNTIL ALL DATA IS READ
	go sendDataLoop(pw, data)

	// DATA IS NOT IN MEMORY
	// YOU WILL READ IT OUT A LITTLE AT A TIME

	// RECEIVE SOME OF MESSAGE
	// WILL READ FROM SHARED MEMORY
	// UNTIL EOF
	for {
		rcvData, err := rcvData(pr)
		fmt.Printf("TEST    : %v \n", err)

		if err != nil {
			fmt.Printf("error with rcvData method: %v", err)
		}
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("RECEIVED    : %s \n", rcvData)
	}

	fmt.Println("DONE")

}
