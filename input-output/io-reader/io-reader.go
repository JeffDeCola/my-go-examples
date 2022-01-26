package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readerToBuffer(r io.Reader) error {

	// CREATE BUFFER TO WRITE TO
	buffer := make([]byte, 20)

	// LETS STREAM FROM READER INTO BUFFER AND PRINT
	for {
		// RETURNS COUNT TILL EOF
		n, err := r.Read(buffer)
		if err == io.EOF {
			return nil
		}
		// A REAL ERROR
		if err != nil {
			return fmt.Errorf("error reading Stdin: %w", err)
		}
		if string(buffer[:4]) == "stop" {
			return nil
		}
		fmt.Printf(string(buffer[:n]) + "\n")
	}

}

func main() {

	// READING FROM A BUFFER - which is a string of bytes
	fmt.Printf("\nREADING FROM A BUFFER\n")
	// CREATE THE READER
	sourceBuffer := []byte("This data is being put into a buffer")
	r := strings.NewReader(string(sourceBuffer))
	// READER TO BUFFER
	err := readerToBuffer(r)
	if err != nil {
		fmt.Printf("Error with reading buffer: %s\n", err)
	}

	// -- READING FROM A FILE (BYTES) TO A BUFFER (BYTES) --
	fmt.Printf("\nREAD FROM FILE\n")
	// OPEN THE FILE
	filename := "test.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %s", err)
	}
	defer f.Close()
	// READER TO BUFFER
	err = readerToBuffer(f)
	if err != nil {
		fmt.Printf("Error with reading file: %s\n", err)
	}

	// -- READING FROM STDIN TO A BUFFER (BYTES) --
	fmt.Printf("\nREAD FROM STDIN - TYPE 'STOP' TO CONTINUE\n")
	// READER TO BUFFER
	err = readerToBuffer(os.Stdin)
	if err != nil {
		fmt.Printf("Error with reading Stdin: %s\n", err)
	}

	// -- READING FROM A PIPE TO A BUFFER (BYTES) --
	fmt.Printf("\nREAD FROM A PIPE\n")
	// CREATE THE PIPE
	rpipe, wpipe := io.Pipe()
	// WRITE INTO PIPE USING GO ROUTINE
	go func() {
		fmt.Fprint(wpipe, "This data is being put into a pipe")
		// Using Close method to close write
		wpipe.Close()
	}()
	// READER TO BUFFER
	err = readerToBuffer(rpipe)
	if err != nil {
		fmt.Printf("Error with reading Pipe: %s\n", err)
	}

}
