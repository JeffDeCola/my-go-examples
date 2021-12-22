# raspi-gpio example

_Using googles `periph` (peripherals in go) library to control a Raspberry Pi's
GPIO (General Purpose Input/Output) pins. The example used will be turning
on/off an external LED via a button._

Documents and references,

* [periph](https://periph.io/)
  home page
* [google/periph](https://github.com/google/periph)
  github page

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## PREREQUISITE

```go
go get -u periph.io/x/periph/cmd/...
```

Will use include as,

```go
    "periph.io/x/periph/conn/gpio"
    "periph.io/x/periph/conn/gpio/gpioreg"
    "periph.io/x/periph/host"
```

I also created following group,

```bash
sudo groupadd gpio
sudo usermod -a -G gpio jeff
sudo grep gpio /etc/group
sudo chown root.gpio /dev/gpiomem
sudo chmod g+rw /dev/gpiomem
```

## GPIO TO BREADBOARD SCHEMATIC

We will control the LED via a button on the Raspberry Pi with
the following setup,

![IMAGE - raspberry-pi-input-and-output-using-gpio-pins - IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/docs/pics/raspberry-pi-input-and-output-using-gpio-pins.jpg?raw=true)

## RUN

```go
go run raspi-gpio.go
```
