package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const toolVersion = "1.0.0"

var errLogLevel = errors.New("please use trace, info or error")

func setLogLevel(logLevel string) error {

	// SET LOG LEVEL (trace, info or error) None of which exit
	log.Trace("Set Log Level")

	switch logLevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.ErrorLevel)
		return fmt.Errorf("%s", errLogLevel)
	}

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	return nil

}

func getNumbers(r1 io.Reader, r2 io.Reader) (int, int, error) {

	// FIRST NUMBER - USER INPUT
	log.Trace("Get 1st number from user and convert to int")
	n1String, err := getUserInput(r1, "What is first number?: ")
	if err != nil {
		return -1, -1, fmt.Errorf("unable to get string from user: %w", err)
	}
	n1, err := convertStringToInt(n1String)
	if err != nil {
		return -1, -1, fmt.Errorf("unable to get string from user: %w", err)
	}

	// SECOND NUMBER - USER INPUT
	log.Trace("Get 2nd number from user and convert to int")
	n2String, err := getUserInput(r2, "What is second number?: ")
	if err != nil {
		return -1, -1, fmt.Errorf("unable to get string from user: %w", err)
	}
	n2, err := convertStringToInt(n2String)
	if err != nil {
		return -1, -1, fmt.Errorf("unable to get string from user: %w", err)
	}

	return n1, n2, nil

}

func getUserInput(r io.Reader, askUser string) (string, error) {

	var nString string

	// GET STRING FROM USER
	log.Trace("Get string from user")
	fmt.Printf("%s", askUser)
	_, err := fmt.Fscan(r, &nString)
	if err != nil {
		return "", fmt.Errorf("unable to get string from user: %w", err)
	}

	return nString, nil

}

func convertStringToInt(nString string) (int, error) {

	// CONVERT STRING TO INT
	log.Trace("Convert string to int")
	n, err := strconv.Atoi(nString)
	if err != nil {
		return -1, fmt.Errorf("unable to convert string to int: %w", err)
	}

	return n, err

}

func getSum(n1 int, n2 int) int {

	// ADD TWO INTS TOGETHER
	log.Trace("Add two integers together")
	result := n1 + n2

	return result

}

func main() {

	// FLAGS
	version := flag.Bool("v", false, "prints current version")
	logLevel := flag.String("loglevel", "error", "log level (trace, info or error)")
	flag.Parse()

	// CHECK VERSION
	if *version {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

	// SET LOG LEVEL (trace, info or error) None of which exit
	err := setLogLevel(*logLevel)
	if err != nil {
		log.Errorf("Error getting logLevel: %s", err)
	}

	// PRINT OUT FLAGS
	log.Trace("Version flag = ", *version)
	log.Trace("Log Level = ", *logLevel)

	fmt.Println(" ")
	fmt.Println("Let's add two numbers together")

	// GET NUMBERS FROM USER INPUT
	n1, n2, err := getNumbers(os.Stdin, os.Stdin)
	if err != nil {
		log.Fatalf("Error getting numbers: %s", err)
	}

	// ADD NUMBERS TOGETHER
	sum := getSum(n1, n2)

	fmt.Printf("Sum is %d\n", sum)
	fmt.Println(" ")

}
