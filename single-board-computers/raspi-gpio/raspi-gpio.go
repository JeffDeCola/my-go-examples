package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	errors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const toolVersion = "1.0.0"

// A Function that also returns an error
func jeffFunc(x int) (string, error) {

	fmt.Println("opening gpio")

	err := rpio.Open()
	if err != nil {
		return "error", errors.New("unable to open gpio")
	}

	defer rpio.Close()

	pin := rpio.Pin(21)
	pin.Output()

	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
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
	r, err := jeffFunc(*integerPtr)
	checkErr(err)
	fmt.Println("Returned", r)

}
