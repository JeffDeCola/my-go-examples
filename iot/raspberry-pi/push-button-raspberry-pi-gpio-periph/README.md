# PUSH BUTTON RASPBERRY PI GPIO PERIPH

_GPIO INPUT - 
Push a button
via a Raspberry Pi GPIO
using the `periph.io/...` packages._

Table of Contents

* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/push-button-raspberry-pi-gpio-periph#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/push-button-raspberry-pi-gpio-periph#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/push-button-raspberry-pi-gpio-periph#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/push-button-raspberry-pi-gpio-periph#test)
* [ILLUSTRATION](https://github.com/JeffDeCola/my-go-examples/tree/master/iot/raspberry-pi/push-button-raspberry-pi-gpio-periph#illustration)

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
* For the Button
  * A 10K Ohm 1/4W resistor (10,000 color code is Brown Black Orange)
  * A 1K Ohm 1/4W resistor (1,000 color code is Brown Black Red)

Where,

* GPIO PIN 23 is connected to the button

To get the entire periph.io go packages,

```bash
go get periph.io/x/cmd/...
```

## OVERVIEW

These are the basic 3 building blocks.

First you init the host machine (i.e. Raspberry Pi),

```
host.Init()
```

Then you assign a GPIO Pin to the button

```go
buttonPin := gpioreg.ByName("23")
```

Since it is an input, we want to set a pulldown resistor and the
ability to look at both edges (High->Low or Low->High) for this pin

```go
buttonPin.In(gpio.PullDown, gpio.BothEdges)
```

Now you can wait for a change in the INPUT of that GPIO pin,

```go
buttonPin.WaitForEdge(-1)
buttonPin.Read()
```

## RUN

```go
go run main.go
```

## TEST

To create test files,

```
gotests -w -all main.go
```

To unit test the code,

```
go test -cover ./...
```

## ILLUSTRATION

How to connect the button and LED to the Raspberry Pi,

![IMAGE - raspberry-pi-input-and-output-using-gpio-pins - IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/docs/pics/raspberry-pi-input-and-output-using-gpio-pins.jpg?raw=true)
