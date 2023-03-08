# TURN LED ON OFF RASPBERRY PI GPIO

_Turn an LED on/off with a button via a Raspberry Pi GPIO
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

Physically, you need,

* Raspberry Pi
* A 560 Ohm 1/4W resistor (560 color code is Green Blue Black)
* An LED
* A button
* A breadboard and some jumper wires

Where,

* GPIO PIN 7 is connected to button
* GPIO PIN 12 to connected to LED

Get the periph.io go packages,

```bash
go get periph.io/x/cmd/...
```

## OVERVIEW

tbd

## EXAMPLE

The code is as follow,

```go
tbd
```

## RUN

You may need to,

```go
sudo groupadd gpio
sudo usermod -a -G gpio jeff
sudo grep gpio /etc/group
sudo chown root.gpio /dev/gpiomem
sudo chmod g+rw /dev/gpiomem
```

```go
go run main.go
```

## TEST



## ILLUSTRATION

![IMAGE - raspberry-pi-input-and-output-using-gpio-pins - IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/docs/pics/raspberry-pi-input-and-output-using-gpio-pins.jpg?raw=true)
