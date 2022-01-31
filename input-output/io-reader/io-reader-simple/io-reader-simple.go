package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	// CREATE THE IO READER INTERFACE
	var r io.Reader

	// CREATE THE BUFFER TO READ FROM
	b := new(bytes.Buffer)
	b.WriteString("This data is being put into a byte.buffer")
	fmt.Printf("Buffer in:  %s\n", b.String())

	// CREATE BUFFER TO WRITE INTO
	buffer := make([]byte, 100)

	// ASSIGN BUFFER TO READER
	r = b

	// READ METHOD (USING io.reader)
	n, err := r.Read(buffer)
	if err != nil {
		fmt.Printf("error reading io.Reader: %s", err)
	}

	// PRINT BUFFER
	fmt.Printf("Buffer out: %s\n", string(buffer[:n]))

}
