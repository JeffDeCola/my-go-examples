package main

import (
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
)

func main() {

	host.Init()

	t := time.NewTicker(500 * time.Millisecond)

	for l := gpio.Low; ; l = !l {
		rpi.P1_33.Out(l)
		<-t.C

	}
}
