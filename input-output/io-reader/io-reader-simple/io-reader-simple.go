package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	// CREATE THE IO READER INTERFACE
	var r io.Reader

	// CREATE THE BUFFER b TO READ FROM
	b := new(bytes.Buffer)
	b.WriteString("This is data in byte.buffer")
	fmt.Printf("Buffer in:  %s\n", b.String())

	// CREATE BUFFER TO WRITE TO
	buffer := make([]byte, 100)

	// ASSIGN BUFFER TO READER
	r = b

	// READ METHOD (USING io.Reader)
	n, err := r.Read(buffer)
	if err != nil {
		fmt.Printf("Error with io.Reader: %s", err)
	}

	// PRINT BUFFER
	fmt.Printf("Buffer out: %s\n", string(buffer[:n]))

}
