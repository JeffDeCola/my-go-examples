// my-go-examples logrus.go

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

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
		return errors.New("please use trace, info or error")
	}

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	return nil

}

func createLogFile(filename string) (*os.File, error) {

	var myfile *os.File

	if filename != "" {

		log.Info("Writing to log file")

		// LOGGING TO FILE (APPEND) (io.Writer)
		myfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(myfile)
		} else {
			return nil, fmt.Errorf("failed to log to file, using default stderr: %w", err)

		}

	}

	return myfile, nil

}

func main() {

	logLevelPtr := flag.String("loglevel", "error", "log level (trace, info or error)")
	logFilePtr := flag.String("logfile", "", "log output to a file")
	flag.Parse()

	// SET LOG LEVEL (trace, info or error) None of which exit
	err := setLogLevel(*logLevelPtr)
	if err != nil {
		log.Errorf("Error getting logLevel: %s", err)
	}

	myfile, err := createLogFile(*logFilePtr)
	if err != nil {
		log.Errorf("Error creating logfile: %s", err)
	}
	// don't forget to close it
	defer myfile.Close()

	// NORMAL LOGGING
	log.Error("Something failed but I'm not quitting.")
	log.Info("Something noteworthy happened!")
	log.Trace("Something very low level.")

	// NORMAL LOGGING WITH FORMATTING
	name := "jeff"
	s := fmt.Sprintf("This is from %s", name)
	log.Info(s)

	// USING FIELDS
	log.WithFields(log.Fields{
		"animal": "cat",
	}).Trace("What animal is it?")

	// REUSING FIELDS
	jeffLogger := log.WithFields(log.Fields{
		"animal": "cat",
		"color":  "grey",
	})

	jeffLogger.Trace("Using the animal and color field")
	jeffLogger.Trace("Me too")

}
