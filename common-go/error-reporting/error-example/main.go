// my-go-examples error-example

package main

import (
	"errors"
	"fmt"

	logger "github.com/JeffDeCola/my-go-packages/golang/logger"
)

// Error definitions
var ErrFilenameEmpty = errors.New("filename can not be empty")

// PROPAGATE ERROR
func firstLevel(filename string) error {

	err := secondLevel(filename)
	if err != nil {
		return fmt.Errorf("firstLevel: %w", err)
	}

	return nil

}

// PROPAGATE ERROR
func secondLevel(filename string) error {

	err := thirdLevel(filename)
	if err != nil {
		return fmt.Errorf("secondLevel: %w", err)
	}

	return nil

}

// ORIGINATE ERROR
func thirdLevel(filename string) error {

	if filename == "" {
		return fmt.Errorf("thirdLevel: %w", ErrFilenameEmpty)
	}

	return nil

}

func main() {

	// Use my logger
	log := logger.CreateLogger(logger.Debug, "jeffs_noTime")

	// Trigger and error by passing empty filename
	err := firstLevel("")
	if err != nil {
		log.Error("Got an Error", "error", err)
	}

	fmt.Println("Optional")

	// (Optional) Unwrap the error level by level
	var chain error = err
	for chain != nil {
		log.Error("Caused by:", "chain", chain)
		chain = errors.Unwrap(chain)
	}

	// (Optional) Unwrap the error until you get root
	chain = err
	for {
		unwrapped := errors.Unwrap(chain)
		if unwrapped == nil {
			log.Error("Root Cause:", "root", chain)
			break
		}
		chain = unwrapped
	}

}
