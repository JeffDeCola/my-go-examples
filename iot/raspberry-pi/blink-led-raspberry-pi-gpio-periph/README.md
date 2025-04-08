# BLINK LED RASPBERRY PI GPIO PERIPH

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_GPIO OUTPUT -
Blink an LED via a Raspberry Pi GPIO using the
[periph.io/...](https://pkg.go.dev/periph.io/x/conn/v3)
packages._

Table of Contents

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/blink-led-raspberry-pi-gpio-periph#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/blink-led-raspberry-pi-gpio-periph#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/blink-led-raspberry-pi-gpio-periph#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/blink-led-raspberry-pi-gpio-periph#test)
* [ILLUSTRATION](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/blink-led-raspberry-pi-gpio-periph#illustration)

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

Where,

* GPIO PIN 5 to connected to the LED

To get the entire periph.io go packages,

```bash
go get periph.io/x/cmd/...
```

## OVERVIEW

These are the basic 3 building blocks.

First you init the host machine (i.e. Raspberry Pi),

```go
host.Init()
```

Then you assign a GPIO Pin to the LED

```go
ledPin := gpioreg.ByName("5")
```

Now you can control the OUTPUT of that GPIO pin
with a High or Low state,

```go
ledPinState := gpio.High
ledPin.Out(ledPinState)
```

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
