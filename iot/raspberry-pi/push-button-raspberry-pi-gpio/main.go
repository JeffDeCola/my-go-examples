// my-go-examples push-button-raspberry-pi-gpio

package main

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func main() {

	// INIT HOST
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// ASSIGN PINS
	button_in := gpioreg.ByName("12")
	if button_in == nil {
		log.Fatal("Failed to find button_in")
	}
	led_out := gpioreg.ByName("5")
	if led_out == nil {
		log.Fatal("Failed to find led_out")
	}

	// SET PULLDOWN FOR BUTTON
	err = button_in.In(gpio.PullDown, gpio.BothEdges)
	if err != nil {
		log.Fatal(err)
	}

	// INIT LED TO ON (true)
	err = led_out.Out(true)
	if err != nil {
		log.Fatal(err)
	}

	toggle := gpio.Low
	count := 0

	for {

		fmt.Printf("  BUTTON is %v\n", button_in.Read())
		fmt.Printf("  LED    is %v\n", led_out.Read())

		// WAIT FOR BUTTON PRESS - RISING EDGE (LOW TO HIGH) - THIS IS BLOCKING
		// -1 disables the timeout
		fmt.Printf("%v Waiting for button press (LOW TO HIGH)...\n", count)
		button_in.WaitForEdge(-1)
		button_in.WaitForEdge(-1)

		// TOOGLE THE LED
		err := led_out.Out(toggle)
		if err != nil {
			log.Fatal(err)
		}

		// time.Sleep(1000 * time.Millisecond)

		// TOGGLE gpio.Low to gpio.High
		toggle = !toggle
		count++

	}

}
