package main

import (
	"fmt"
	"io"
	"strings"
)

func sendData(pw *io.PipeWriter, data string) {

	_, err := pw.Write([]byte(data))
	if err != nil {
		fmt.Printf("error with pr.Write method: %s", err)
	}

}

func receiveData(pr *io.PipeReader) (string, error) {

	buffer := make([]byte, 100)

	_, err := pr.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("error with pr.Read method: %w", err)
	}

	// REPLACE WITH "that will be sent" with "that was received"
	replacedString := strings.Replace(string(buffer), "that will be sent", "that was received", -1)

	return replacedString, nil
}

func main() {

	pr, pw := io.Pipe()

	// DATA TO SEND
	dataSend := "I am the data that will be sent"
	fmt.Printf("SEND:     %s \n", dataSend)

	// SEND
	go sendData(pw, dataSend)

	// RECEIVE
	dataReceived, err := receiveData(pr)
	if err != nil {
		fmt.Printf("Error with receiving data: %s", err)
	}

	// PRINT
	fmt.Printf("RECEIVED: %s \n", dataReceived)

}
