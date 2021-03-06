package main

import (
	"flag"
	"fmt"
	"os"

	errors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const toolVersion = "1.0.1"

var integerPtr *int

// A Function that also returns an error
func jeffFunc(x int) (string, error) {
	if x == 42 {
		// Make your error
		return "error", errors.New("can't work with 42")
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
	// INTEGER FLAG
	integerPtr = flag.Int("i", 42, "This is the flag for an integer")
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
