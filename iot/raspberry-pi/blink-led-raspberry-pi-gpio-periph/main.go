// my-go-examples blink-led-raspberry-pi-gpio-periph

package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func blinkLed(ledPin gpio.PinIO, blinkTime time.Duration) {

	ledPinState := gpio.High
	ticker := time.NewTicker(blinkTime * time.Millisecond)

	for range ticker.C {

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

func main() {

	// HOW FAST TO BLINK
	var blinkTime time.Duration = 500

	// INIT HOST MACHINE (i.e. Raspberry Pi)
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// ASSIGN GPIO5 TO ledPin
	ledPin := gpioreg.ByName("5")
	if ledPin == nil {
		log.Fatal("Failed to find ledPin")
	}

	// KICK OFF BLINK LOOP
	go blinkLed(ledPin, blinkTime)

	// PRESS ENTER TO EXIT
	fmt.Println("Press Enter to Stop")
	fmt.Scanln()

}
