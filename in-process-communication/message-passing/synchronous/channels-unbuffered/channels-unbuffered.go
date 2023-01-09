package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func sendData(chData chan string, chUnBuffered chan string) {

	// Loop forever
	for {

		// Data to send
		data := <-chData

		// Send data to buffered channel
		chUnBuffered <- data // BLOCKING - WAITING
	}

}

func rcvData(chUnBuffered chan string) {

	// Loop forever every 2 seconds
	for {

		data := <-chUnBuffered // BLOCKS

		// REPLACE WITH "that will be sent" with "that was received"
		replacedString := strings.Replace(data, "that will be sent", "that was received", -1)

		//REMOVE TIME AT END
		past, _ := strconv.Atoi(replacedString[strings.LastIndex(replacedString, " ")+1:])
		now, _ := strconv.Atoi(strconv.FormatInt(time.Now().Unix(), 10))
		diff := now - past
		// REMOVE EPOCH TIME
		replacedString = strings.Replace(replacedString, strconv.Itoa(past), "", -1)

		// PRINT
		fmt.Printf("RECEIVED: %s- Took %s seconds \n", replacedString, strconv.Itoa(diff))

		// TIME TO DO THINGS
		time.Sleep(2 * time.Second) // Just chill

	}

}

func main() {

	chData := make(chan string)
	chUnBuffered := make(chan string) // UNBUFFERED

	// DATA TO SEND
	data := "I am the data that will be sent"
	fmt.Printf("SEND:     %s \n", data)

	// SEND GO ROUTINE
	go sendData(chData, chUnBuffered)

	// RECEIVE GO ROUTINE - Can consume every 2 seconds
	go rcvData(chUnBuffered)

	startTime := time.Now().Unix()

	// SEND DATA WITH EPOCH TIME - FAST
	fmt.Println("SEND FAST - WAITS")
	chData <- data + " 1 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chData <- data + " 2 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chData <- data + " 3 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chData <- data + " 4 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chData <- data + " 5 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chData <- data + " 6 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(5 * time.Second)

	// SEND DATA WITH EPOCH TIME - SLOW
	fmt.Println("SEND SLOW")
	chData <- data + " A " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(3 * time.Second)
	chData <- data + " B " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(3 * time.Second)
	chData <- data + " C " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(3 * time.Second)
	chData <- data + " D " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(3 * time.Second)
	chData <- data + " E " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS

	endTime := time.Now().Unix()
	diffTime := endTime - startTime
	fmt.Printf("TOTAL TIME: %d seconds\n", diffTime)

	time.Sleep(5 * time.Second)

}
