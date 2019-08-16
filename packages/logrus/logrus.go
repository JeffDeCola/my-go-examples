// my-go-examples logrus.go

package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {

	// SET LOG LEVEL
	// log.SetLevel(log.InfoLevel)
	log.SetLevel(log.TraceLevel)

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	// LOGGING TO FILE (APPEND) (io.Writer)
	/* myfile, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(myfile)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	// don't forget to close it
	defer myfile.Close() */

	// NORMAL LOGGING
	//log.Panic("I'm bailing. Calls panic() after logging.")
	//log.Fatal("Bye. Calls os.Exit(1) after logging.")
	log.Error("Something failed but I'm not quitting.")
	log.Warn("You should probably take a look at this.")
	log.Info("Something noteworthy happened!")
	log.Debug("Useful debugging information.")
	log.Trace("Something very low level.")

	// NORMAL LOGGING WITH FORMATTING
	name := "jeff"
	s := fmt.Sprintf("This is from %s", name)
	log.Info(s)

	// USING FIELDS
	log.WithFields(log.Fields{
		"animal": "cat",
	}).Info("A cat appears")

	// REUSING FIELDS
	jeffLogger := log.WithFields(log.Fields{
		"animal": "cat",
		"other":  "I also should be logged always",
	})
	jeffLogger.Info("I'll be logged with common and other field")
	jeffLogger.Info("Me too")

}
