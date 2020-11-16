package main

import (
    "fmt"
    "log"
    "time"

    "periph.io/x/periph/conn/gpio"
    "periph.io/x/periph/host"
    "periph.io/x/periph/host/rpi"
)

func main() {
    // Load all the drivers:
    if _, err := host.Init(); err != nil {
        log.Fatal(err)
    }

    // Lookup a pin by its number:
    p20 := gpioreg.ByName("GPI20")
    if p20 == nil {
        log.Fatal("Failed to find GPI20")
    }

    fmt.Printf("%s: %s\n", p20, p20.Function())

    // Set it as input, with an internal pull down resistor:
    if err = p20.In(gpio.PullDown, gpio.BothEdges); err != nil {
        log.Fatal(err)
    }

    // Wait for edges as detected by the hardware, and print the value read:
    for {
        p20.WaitForEdge(-1)
        fmt.Printf("-> %s\n", p20.Read())
    }

    t := time.NewTicker(500 * time.Millisecond)
    for l := gpio.Low; ; l = !l {
        // Lookup a pin by its location on the board:
        if err := rpi.P1_40.Out(l); err != nil {
            log.Fatal(err)
        }
        <-t.C
    }
}

