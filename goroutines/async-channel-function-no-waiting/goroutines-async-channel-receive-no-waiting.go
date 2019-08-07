package main

import (
	"fmt"
	"strconv"
	"time"
)

var sendSpeedSeconds = 10
var rcvSpeedSeconds = 2

// sendData - Sends data every X seconds
func sendData(sendCh chan string) {

	var data string
	counter := 1

	// Loop forever - Long Running (Will start immediately then wait for tick)
	for c := time.Tick(time.Duration(sendSpeedSeconds) * time.Second); ; <-c {

		fmt.Printf("%4d - START SENDING DATA\n", counter)

		// Send Data
		data = strconv.Itoa(counter)
		sendCh <- data
		fmt.Printf("%4d - Sent data %v\n", counter, data)
		counter++
	}
}

// rcvData - Checks for data every x seconds - Does not wait.
func rcvData(rcvCh chan string) {

	data := "empty"
	counter := 1

	// Loop forever - Long Running (Will start immediately then wait for tick)
	for c := time.Tick(time.Duration(rcvSpeedSeconds) * time.Second); ; <-c {

		fmt.Printf("%40d - START RECEIVING DATA\n", counter)

		// Get the current data list (from channel)

		// If there is something in channel
		//       - Read and continue reading until empty channel (hence, get latest one)
		//       - break out of loop like above
		// If there is nothing in channel
		//       - default and break out of the loop
		for {
			select {
			case newVal := <-rcvCh:
				data = newVal
				fmt.Printf("%40d - Received data %v\n", counter, data)
				continue
			default:
			}
			break
		}

		fmt.Printf("%40d - Using data %v\n", counter, data)
		counter++
	}
}

func main() {

	fmt.Println(" ")
	fmt.Printf("sendData() will send data every %v seconds\n", sendSpeedSeconds)
	fmt.Printf("rcvData() will receive data every %v seconds\n", rcvSpeedSeconds)
	fmt.Println(" ")

	var msgCh = make(chan string, 100)
	go sendData(msgCh)
	go rcvData(msgCh)

	// Press return to exit
	fmt.Scanln()
	fmt.Println("done")
}
