//my-go-examples

//my-go-examples

package main

import (
	"fmt"
	"os"

	logger "github.com/JeffDeCola/my-go-packages/golang/logger"
)

func main() {

	// Open or create the log file, output is filename
	output, err := os.OpenFile("jefflog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open/create log file" + err.Error())
	}

	log := logger.CreateLogger(logger.Trace, "jeffs_noTime", output)

	a := 4.54534

	log.Trace("This is a trace message")
	log.Debug("This is a debug message")
	log.Info(fmt.Sprintf("Formatted Info Message a=%.2f", a), "a", a, "user", "jeff")
	log.Warning("This is a Warning Message", "user", "jeff")
	log.Error("This is an Error message")
	// log.Fatal("Fatal Error")

	// Dynamically change log level
	fmt.Printf("\nCHANGE LEVEL\n\n")
	log.ChangeLogLevel(logger.Warning)

	log.Trace("This is a trace message")
	log.Debug("This is a debug message")
	log.Info(fmt.Sprintf("Formatted Info Message a=%.2f", a), "a", a, "user", "jeff")
	log.Warning("This is a Warning Message", "user", "jeff")
	log.Error("This is an Error message")
}
