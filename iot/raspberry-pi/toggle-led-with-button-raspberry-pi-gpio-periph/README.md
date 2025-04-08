# TOGGLE LED WITH BUTTON RASPBERRY PI GPIO PERIPH

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Toggle an LED with a button push via a Raspberry Pi GPIO using the
[periph.io/...](https://pkg.go.dev/periph.io/x/conn/v3)
packages._

Table of Contents

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/toggle-led-with-button-raspberry-pi-gpio-periph#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/toggle-led-with-button-raspberry-pi-gpio-periph#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/toggle-led-with-button-raspberry-pi-gpio-periph#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/toggle-led-with-button-raspberry-pi-gpio-periph#test)
* [ILLUSTRATION](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/toggle-led-with-button-raspberry-pi-gpio-periph#illustration)

Documentation and Reference

* [periph.io/x/conn/v3/gpio](https://pkg.go.dev/periph.io/x/conn/v3/gpio)
  package defines the digital pins
* [periph.io/x/host/v3/rpi](https://pkg.go.dev/periph.io/x/host/v3/rpi)
  package defines the Raspberry Pi
* Raspberry Pi
  [install and configure](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet)
* Raspberry Pi
  [specifications](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/raspberry-pi/specifications-cheat-sheet)

## PREREQUISITES

Physically, you will need,

* Raspberry Pi
* Breadboard with a cable connector and jumper wires
* For the LED
  * A 560 Ohm 1/4W resistor (560 color code is Green Blue Black)
* For the Button
  * A 10K Ohm 1/4W resistor (10,000 color code is Brown Black Orange)
  * A 1K Ohm 1/4W resistor (1,000 color code is Brown Black Red)

Where,

* GPIO PIN 23 is connected to the button
* GPIO PIN 5 to connected to the LED

To get the entire periph.io go packages,

```bash
go get periph.io/x/cmd/...
```

## OVERVIEW

The program will kick off a button and a LED go routine.

The button go routine is shown in my
[push-button-raspberry-pi-gpio-periph](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/push-button-raspberry-pi-gpio-periph)
INPUT example.

The LED go routine was shown in my
[blink-led-raspberry-pi-gpio-periph](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/blink-led-raspberry-pi-gpio-periph)
OUTPUT example.

I connect bother together using a channel.  When the
button is fully pressed I send this "event"
from the  button go routine
to the LED go routine
to toggle the LED.

## RUN

```go
go run main.go
```

## TEST

To create test files,

```bash
gotests -w -all main.go
```

To unit test the code,

```bash
go test -cover ./...
```

## ILLUSTRATION

How to connect the button and LED to the Raspberry Pi,

![IMAGE - raspberry-pi-input-and-output-using-gpio-pins - IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/docs/pics/other/raspberry-pi-input-and-output-using-gpio-pins.svg?raw=true)
