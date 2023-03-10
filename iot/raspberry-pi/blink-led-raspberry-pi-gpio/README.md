# BLINK LED RASPBERRY PI GPIO

  _Blink an LED
  via a Raspberry Pi GPIO
  using the `periph.io/...` packages._

Table of Contents

* tbd

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

First you init the host machine (i.e. Raspberry {Pi),

```
host.Init()
```

Then you assign a GPIO Pin to the LED

```go
ledPin := gpioreg.ByName("5")
```

Now you can control that GPIO pin with a High or Low state,

```go
ledPinState := gpio.High 
ledPin.Out(ledPinState)
```

## RUN

```go
go run main.go
```

## TEST

To create _test files,

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
