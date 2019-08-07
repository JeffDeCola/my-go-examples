package main

import (
	"fmt"
	"time"
)

func talk(msgCh chan string) {

	count := 0
	say := "empty"

	// Loop forever - Long Running (Will start immediately then wait for tick)
	for {
		// time.Sleep(5 * time.Second)
		say = <-msgCh // WAITS
		time.Sleep(2 * time.Millisecond)
		fmt.Printf("%30d Received message %s\n", count, say)
		count++
	}
}

func main() {

	msgCh := make(chan string)

	// A goroutine calling a method with a channel for communication
	go talk(msgCh)

	time.Sleep(5 * time.Second)
	msgCh <- "Jeff" // WAITS
	fmt.Printf("Sent message %s\n", "Jeff")
	msgCh <- "Clif" // WAITS
	fmt.Printf("Sent message %s\n", "Clif")

	time.Sleep(8 * time.Second)
	msgCh <- "Jack" // WAITS
	fmt.Printf("Sent message %s\n", "Jack")
	msgCh <- "Jill" // WAITS
	fmt.Printf("Sent message %s\n", "Jill")

	time.Sleep(4 * time.Second)
	msgCh <- "quit"
}
