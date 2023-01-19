package main

import (
	"fmt"
	"io"
)

func main() {

	// CREATE UNNAMED PIPE
	pr, pw := io.Pipe()

	// DATA TO SEND
	data := "I am the data that will be sent"
	fmt.Printf("SEND       : %s \n", data)

	// SEND - WRITE TO SHARED MEMORY
	go func() {
		defer pw.Close()
		_, err := pw.Write([]byte(data))
		if err != nil {
			fmt.Printf("error with pr.Write method: %v", err)
		}
	}()

	// RECEIVE - READ FROM SHARED MEMORY
	rcvData := make([]byte, 100)
	_, err := pr.Read(rcvData)
	if err != nil {
		fmt.Printf("error with pr.Read method: %v", err)
	}
	fmt.Printf("RECEIVED   : %s \n", rcvData)

}
