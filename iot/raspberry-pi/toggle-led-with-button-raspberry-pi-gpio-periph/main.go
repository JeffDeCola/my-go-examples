// my-go-examples toggle-led-with-button-raspberry-pi-gpio-periph

package main

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func toggleLED(ledPin gpio.PinIO, myIntChannel chan bool) {

	ledPinState := gpio.High

	for {

		// WAITS - BLOCKS (DON'T CARE ABOUT DATA)
		<-myIntChannel

		// PIN OUTPUT FOR LED ON TOGGLE VALUE
		err := ledPin.Out(ledPinState)
		if err != nil {
			log.Fatal(err)
		}

		// LED PIN STATUS (High or low)
		fmt.Printf("LED is %v\n", ledPin.Read())

		// TOGGLE PIN OUTPUT FOR LED
		ledPinState = !ledPinState

	}

}

func buttonStateChange(buttonPin gpio.PinIO, myIntChannel chan bool) {

	count := 1
	var buttonPinLevel gpio.Level
	state := "low"

	// THE BUTTON MUST GO FROM LOW TO HIGH TO LOW
	for {

		// WAIT FOR BUTTON PRESS - RISING EDGE (LOW TO HIGH) - THIS IS BLOCKING
		// -1 disables the timeout
		buttonPin.WaitForEdge(-1)

		// GET BUTTON LEVEL
		buttonPinLevel = buttonPin.Read()

		switch state {
		case "low":
			if buttonPinLevel == gpio.High {
				state = "high"
			}
		case "high":
			if buttonPinLevel == gpio.Low {
				state = "low"
				fmt.Printf("Pressed Button: #%4d\n", count)
				myIntChannel <- true // Really doesn't matter, Could be false
				count++
			}

		}

	}

}

func main() {

	myIntChannel := make(chan bool)

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

	// ASSIGN GPIO5 TO ledPin
	ledPin := gpioreg.ByName("5")
	if ledPin == nil {
		log.Fatal("Failed to find ledPin")
	}

	// SET PULLDOWN RESISTER AND LOOK FOR BOTH EDGES (High->Low or Low->High) FOR THIS PIN
	err = buttonPin.In(gpio.PullDown, gpio.BothEdges)
	if err != nil {
		log.Fatal(err)
	}

	// KICK OFF BUTTON STATE CHANGE LOOP
	go buttonStateChange(buttonPin, myIntChannel)

	// KICK OFF BUTTON STATE CHANGE LOOP
	go toggleLED(ledPin, myIntChannel)

	// PRESS ENTER TO EXIT
	fmt.Println("Press Enter to Stop")
	fmt.Scanln()

}
