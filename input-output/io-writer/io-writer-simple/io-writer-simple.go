package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	// CREATE THE IO WRITER INTERFACE
	var w io.Writer

	// CREATE BUFFER TO READ FROM
	buffer := []byte("This is the buffer data")

	// CREATE THE BUFFER b TO WRITE TO
	b := new(bytes.Buffer)
	fmt.Printf("Buffer in:  %s\n", b.String())

	// ASSIGN b TO WRITER
	w = b

	// WRITE METHOD (USING io.Writer)
	_, err := w.Write(buffer)
	if err != nil {
		fmt.Printf("Error with io.Writer: %s", err)
	}

	// PRINT b
	fmt.Printf("Buffer out: %s\n", b.String())

}
