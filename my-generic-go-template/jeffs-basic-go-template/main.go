package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	logger "github.com/JeffDeCola/my-go-packages/golang/logger"
)

// Create Logger
var log = logger.CreateLogger(logger.Error, "jeffs_noTime", os.Stdout)

const toolVersion = "1.0.0"

type flagConfig struct {
	LogLevel string
	Version  bool
	Help     bool
}

func processFlags() (cfg *flagConfig) {

	// Initialize cfg
	cfg = &flagConfig{}

	// Assign flags to flagConfig Struct
	flag.StringVar(&cfg.LogLevel, "loglevel", "error", "log level (trace, debug, info, error, warning, fatal)")
	flag.BoolVar(&cfg.Version, "v", false, "prints current version")
	flag.BoolVar(&cfg.Help, "h", false, "displays help information")
	flag.Parse()

	// USAGE
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "jeffs-basic-go-template v%s\n\n", toolVersion)
		fmt.Fprintf(os.Stderr, "A simple program that adds two numbers together.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  jeffs-basic-go-template [flags]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  jeffs-basic-go-template\n")
		fmt.Fprintf(os.Stderr, "  jeffs-basic-go-template -loglevel debug\n")
		fmt.Fprintf(os.Stderr, "  jeffs-basic-go-template -v\n")
	}

	if cfg.Help {
		flag.Usage()
		os.Exit(0)
	}

	if cfg.Version {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

	return cfg

}

func main() {

	cfg := processFlags()
	fmt.Println(cfg)

	// SET LOG LEVEL
	logLevel := logger.MyLogLevel(0)
	log.ChangeLogLevel(logLevel)

	fmt.Println(" ")
	fmt.Println("Let's add two numbers together")

	// GET NUMBERS FROM USER INPUT
	n1, n2, err := getNumbers(os.Stdin, os.Stdin)
	if err != nil {
		log.Error("Error getting numbers:", "error", err)
	}

	// ADD NUMBERS TOGETHER
	sum := getSum(n1, n2)

	fmt.Printf("Sum is %d\n", sum)
	fmt.Println(" ")

}

func getNumbers(r1 io.Reader, r2 io.Reader) (n1, n2 int, err error) {

	// FIRST NUMBER - USER INPUT
	log.Trace("Get 1st number from user and convert to int")
	n1String, err := getUserInput(r1, "What is first number?: ")
	if err != nil {
		return -1, -1, fmt.Errorf("getNumbers: %w", err)
	}
	n1, err = convertStringToInt(n1String)
	if err != nil {
		return -1, -1, fmt.Errorf("getNumbers: %w", err)
	}

	// SECOND NUMBER - USER INPUT
	log.Trace("Get 2nd number from user and convert to int")
	n2String, err := getUserInput(r2, "What is second number?: ")
	if err != nil {
		return -1, -1, fmt.Errorf("getNumbers: %w", err)
	}
	n2, err = convertStringToInt(n2String)
	if err != nil {
		return -1, -1, fmt.Errorf("getNumbers: %w", err)
	}

	return n1, n2, nil

}

func getUserInput(r io.Reader, askUser string) (input string, err error) {

	var nString string

	// GET STRING FROM USER
	log.Trace("Get string from user")
	fmt.Printf("%s", askUser)
	_, err = fmt.Fscan(r, &nString)
	if err != nil {
		return "", fmt.Errorf("getUserInput: %w", err)
	}

	return nString, nil

}

func convertStringToInt(nString string) (nInt int, err error) {

	// CONVERT STRING TO INT
	log.Trace("Convert string to int")
	n, err := strconv.Atoi(nString)
	if err != nil {
		return -1, fmt.Errorf("convertStringToInt: %w", err)
	}

	return n, err

}

func getSum(n1 int, n2 int) int {

	// ADD TWO INTS TOGETHER
	log.Trace("Add two integers together")
	result := n1 + n2

	return result

}
