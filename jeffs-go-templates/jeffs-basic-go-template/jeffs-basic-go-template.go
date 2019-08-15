package main

import (
	"fmt"
	"os"

	errors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// A Function that also returns an error
func jeffFunc(x int) (int, error) {
	if x == 42 {
		// Make your error
		return -1, errors.New("can't work with 42")
	}
	return x + 3, nil
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
}

func main() {

	log.Trace("Calling jeffFunc to get result")
	r, err := jeffFunc(43)
	checkErr(err)
	fmt.Println("Returned", r)

}
