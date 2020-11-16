package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"

	log "github.com/sirupsen/logrus"
)

const toolVersion = "1.0.0"

// A Function that also returns an error
func jeffFunc() (string, error) {

	fmt.Println("opening gpio")

	host.Init()
	t := time.NewTicker(500 * time.Millisecond)
	for l := gpio.Low; ; l = !l {
		rpi.P1_33.Out(l)
		<-t.C
	}

	return "Everything worked great", nil
}

// Check your error
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func init() {

	// SET LOG LEVEL
	// log.SetLevel(log.InfoLevel)
	log.SetLevel(log.TraceLevel)

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	// VERSION FLAG
	version := flag.Bool("v", false, "prints current version")
	// Parse the flags
	flag.Parse()

	// CHECK VERSION
	if *version {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

}

func main() {

	log.Trace("Calling jeffFunc to get result")
	r, err := jeffFunc()
	checkErr(err)
	fmt.Println("Returned", r)

}
