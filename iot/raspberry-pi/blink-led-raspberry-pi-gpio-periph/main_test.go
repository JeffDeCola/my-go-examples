// my-go-examples blink-led-raspberry-pi-gpio

package main

import (
	"testing"
	"time"

	"periph.io/x/conn/v3/gpio"
)

func Test_blinkLed(t *testing.T) {
	type args struct {
		ledPin    gpio.PinIO
		blinkTime time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blinkLed(tt.args.ledPin, tt.args.blinkTime)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
