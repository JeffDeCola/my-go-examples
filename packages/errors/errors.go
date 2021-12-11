// my-go-examples errors.go

package main

import (
	"errors"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func a() (string, error) {

	// DO STUFF

	// CALL ANOTHER FUNCTION
	out, err := b()
	if err != nil {
		return out, fmt.Errorf("error in a: %w", err)
	}

	return out, err

}

func b() (string, error) {

	// DO STUFF

	// CALL ANOTHER FUNCTION
	out, err := c()
	if err != nil {
		return out, fmt.Errorf("error in b: %w", err)
	}

	return out, err

}

func c() (string, error) {

	var userAnswer string

	// GET USER ANSWER
	fmt.Print("What is 2 + 2?: ")
	_, err := fmt.Scan(&userAnswer)
	if err != nil {
		return "", fmt.Errorf("unable to get user input: %w", err)
	}

	if userAnswer == "stop" {
		return userAnswer, nil
	}

	// CONVERT TO INT
	answerInt, err := strconv.Atoi(userAnswer)
	if err != nil {
		return userAnswer, fmt.Errorf("unable to convert string: %w", err)
	}

	// IS ANSWER CORRECT?
	if answerInt != 4 {
		// Make your error
		err := errors.New("incorrect answer " + userAnswer + "error in c")
		return userAnswer, err
	}

	return userAnswer, nil

}

func main() {

	// ASK USER A MATH QUESTION UNTIL THEY TYPE "stop"
	for {

		userAnswer, err := a()
		if err != nil {
			log.Errorf("INCORRECT: %s", err)
		}

		if userAnswer == "stop" {
			fmt.Println("Done")
			break
		} else {
			fmt.Println("CORRECT: ", userAnswer)
		}

		fmt.Println("-----------------------------")
	}

}
