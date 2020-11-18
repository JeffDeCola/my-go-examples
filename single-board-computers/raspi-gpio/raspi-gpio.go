// my-go-examples raspi-gpio.go

package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {

	// Load all the drivers: Initialize the periph host
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Lookup pins by their number
	button_in := gpioreg.ByName("GPIO20")
	if button_in == nil {
		log.Fatal("Failed to find GPI20")
	}
	led_out := gpioreg.ByName("GPIO21")
	if led_out == nil {
		log.Fatal("Failed to find GPI21")
	}

	// Setup P20 as INPUT
	// the RasPi has an internal pull down resistor
	// PULL DOWN - DEFAULTS TO LOW WHEN UNCONNECTED
	// RISING EDGE - LOW TO HIGH
	err = button_in.In(gpio.PullDown, gpio.RisingEdge)
	if err != nil {
		log.Fatal(err)
	}
	// Setup P21 as OUTPUT
	// With output level 1 (true)
	err = led_out.Out(true)
	if err != nil {
		log.Fatal(err)
	}

	toggle := gpio.Low
	count := 0

	for {

		fmt.Printf("  Pin 20 (BUTTON) is %s: %s (%s)\n", button_in, button_in.Function(), button_in.Read())
		fmt.Printf("  Pin 21 (LED)    is %s: %s (%s)\n", led_out, led_out.Function(), led_out.Read())

		// WAIT FOR BUTTON PRESS - RISING EDGE (LOW TO HIGH) - THIS IS BLOCKING
		// -1 disables the timeout
		fmt.Printf("%v Waiting for button press (LOW TO HIGH)...\n", count)
		button_in.WaitForEdge(-1)

		// TOOGLE THE LED
		err := led_out.Out(toggle)
		if err != nil {
			log.Fatal(err)
		}

		// Change gpio.Low to gpio.High
		toggle = !toggle
		count++

	}
}
