package main

import (
	"bytes"
	"fmt"
	"io"
)

func writeFromBuffer(w io.Writer) error {

	// WRITE METHOD
	// Returns byte count

	fmt.Printf("Wrote %d bytes to io.Writer\n", n)
	return nil

}

func main() {

	// CREATE BUFFER
	buffer := []byte("This data is being put into a byte reader\n")

	// CREATE THE BUFFER
	b := new(bytes.Buffer)
	fmt.Printf("Buffer before io.Writer: %s \n", b)

	// STREAM FROM BUFFER (b is *bytes.Buffer)
	err := writeFromBuffer(b)
	if err != nil {
		fmt.Printf("Error writing to buffer: %s\n", err)
	}

	n, err := w.Write(buffer)
	if err != nil {
		return fmt.Errorf("error writing io.Writer: %w", err)
	}
	// PRINT BUFFER
	fmt.Printf("Buffer after io.Writer: %s \n", b)

}
