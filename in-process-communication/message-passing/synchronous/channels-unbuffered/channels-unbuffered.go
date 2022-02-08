package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func sendData(chSend chan string, chMsg chan string) {

	data := "empty"

	for {
		data = <-chSend // BLOCKS
		chMsg <- data   // WAITS
	}

}

func rcvData(chMsg chan string) {

	for {

		data := <-chMsg // BLOCKS

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

	chSend := make(chan string)
	chMsg := make(chan string) // UNBUFFERED

	// DATA TO SEND
	dataSend := "I am the data that will be sent"
	fmt.Printf("SEND:     %s \n", dataSend)

	// SEND
	go sendData(chSend, chMsg)

	// RECEIVE - Can consume every 2 seconds
	go rcvData(chMsg)

	startTime := time.Now().Unix()

	// SEND DATA WITH EPOCH TIME - FAST
	fmt.Println("SEND FAST - WAITS")
	chSend <- dataSend + " 1 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chSend <- dataSend + " 2 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chSend <- dataSend + " 3 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chSend <- dataSend + " 4 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chSend <- dataSend + " 5 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	chSend <- dataSend + " 6 " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(5 * time.Second)

	// SEND DATA WITH EPOCH TIME - SLOW
	fmt.Println("SEND SLOW")
	chSend <- dataSend + " A " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(3 * time.Second)
	chSend <- dataSend + " B " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(3 * time.Second)
	chSend <- dataSend + " C " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(3 * time.Second)
	chSend <- dataSend + " D " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS
	time.Sleep(3 * time.Second)
	chSend <- dataSend + " E " + strconv.FormatInt(time.Now().Unix(), 10) // WAITS

	endTime := time.Now().Unix()
	diffTime := endTime - startTime
	fmt.Printf("TOTAL TIME: %d seconds\n", diffTime)

	time.Sleep(5 * time.Second)

}
