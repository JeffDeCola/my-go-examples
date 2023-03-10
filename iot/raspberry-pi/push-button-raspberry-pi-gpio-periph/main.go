// my-go-examples push-button-raspberry-pi-gpio-periph

package main

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func buttonStateChange(buttonPin gpio.PinIO) {

	count := 1

	for {

		fmt.Printf("%4d  BUTTON is %v\n", count, buttonPin.Read())

		// WAIT FOR BUTTON PRESS - RISING EDGE (LOW TO HIGH) - THIS IS BLOCKING
		// -1 disables the timeout
		buttonPin.WaitForEdge(-1)

		count++

	}

}

func main() {

	// INIT HOST MACHINE (i.e. Raspberry Pi)
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// ASSIGN GPIO12 TO buttonPin
	buttonPin := gpioreg.ByName("23")
	if buttonPin == nil {
		log.Fatal("Failed to find buttonPin")
	}

	// SET PULLDOWN RESISTER AND LOOK FOR BOTH EDGES (High->Low or Low->High) FOR THIS PIN
	err = buttonPin.In(gpio.PullDown, gpio.BothEdges)
	if err != nil {
		log.Fatal(err)
	}

	// KICK OFF BUTTON STATE CHANGE LOOP
	go buttonStateChange(buttonPin)

	// PRESS ENTER TO EXIT
	fmt.Println("Press Enter to Stop")
	fmt.Scanln()

}
