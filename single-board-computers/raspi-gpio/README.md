# raspi-gpio example

_Using googles `periph` library to control a Raspberry Pi's
GPIO input/output pins. The example used will be turning
on/off an external LED via a button._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITE

```go
go get -u periph.io/x/periph/cmd/...
```

Will include as,

```go
    "periph.io/x/periph/conn/gpio"
    "periph.io/x/periph/conn/gpio/gpioreg"
    "periph.io/x/periph/host"
```

## OVERVIEW

We will control the LED via a button on the Raspberry Pi as follows,

![IMAGE - raspberry-pi-input-and-output-using-gpio-pins - IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/docs/pics/raspberry-pi-input-and-output-using-gpio-pins.jpg?raw=true)

## RUN

```go
go run raspi-gpio.go
```
