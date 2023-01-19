package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
)

var fifoFile = "/tmp/myFifo"

func writer(filePath string) {
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0600)

	fmt.Printf("WRITER << opened: %+v|%+v\n", f, err)
	if err != nil {
		panic(err)
	}

	fmt.Printf("WRITER << encoder created\n")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		_, err = f.WriteString(fmt.Sprint("line", i, "\n"))
		fmt.Printf("WRITER << written line%d, %+v\n", i, err)
	}

	time.Sleep(1 * time.Second)
	err = f.Close()
	fmt.Printf("WRITER << closed: %+v\n", err)
}

func reader(filePath string) {

	// Delete existing pipes
	fmt.Println("Cleanup existing FIFO file")
	os.Remove(filePath)

	// Create pipe
	fmt.Println("Creating " + filePath + " FIFO file")
	err := syscall.Mkfifo(filePath, 0640)
	if err != nil {
		fmt.Println("Failed to create pipe")
		panic(err)
	}

	// Open pipe for read only
	fmt.Println("Starting read operation")
	pipe, err := os.OpenFile(fifoFile, os.O_RDONLY, 0640)
	if err != nil {
		fmt.Println("Couldn't open pipe with error: ", err)
	}
	defer pipe.Close()

	// Read the content of named pipe
	reader := bufio.NewReader(pipe)
	fmt.Println("READER >> created")

	// Infinite loop
	for {
		line, err := reader.ReadBytes('\n')
		// Close the pipe once EOF is reached
		if err != nil {
			fmt.Println("FINISHED!")
			os.Exit(0)
		}

		// Remove new line char
		nline := string(line)
		nline = strings.TrimRight(nline, "\r\n")
		fmt.Printf("READER >> reading line: %+v\n", nline)

	}

}

func main() {
	fmt.Printf("STARTED %s\n", fifoFile)
	go writer(fifoFile)
	reader(fifoFile)

}
