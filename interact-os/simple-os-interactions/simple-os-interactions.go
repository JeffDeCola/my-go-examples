package main

import (
	"fmt"
	"syscall"
)

func main() {

	// GET PROCESS ID (pid)
	pid := syscall.Getpid()
	fmt.Printf("Process ID (pid): %d\n", pid)

	// GET THREAD ID (tid)
	tid := syscall.Gettid()
	fmt.Printf("Thread ID (id):   %d\n", tid)

}
