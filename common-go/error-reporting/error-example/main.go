// my-go-examples error-example

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// ErrIncorrectAnswer is an error that is returned when the answer is incorrect
var ErrIncorrectAnswer = errors.New("the answer is incorrect")

func askQuestion(r io.Reader) (string, error) {

	var userInput string

	fmt.Print("What is 2+2 (type stop to quit)? ")
	// THIS WILL BLOCK
	_, err := fmt.Fscan(r, &userInput)
	if err != nil {
		return "", fmt.Errorf("unable to get string from user: %w", err)
	}

	return userInput, nil

}

func checkAnswer(answer string) error {

	// CONVERT TO INT
	answerInt, err := strconv.Atoi(answer)
	if err != nil {
		return fmt.Errorf("unable to convert to integer: %w", err)
	}

	// CHECK NUMBER
	err = checkNumber(answerInt)
	if err != nil {
		return fmt.Errorf("error calling checkNumber: %w", err)
	}

	return err

}

func checkNumber(answer int) error {

	// IS ANSWER CORRECT?
	if answer != 4 {
		return fmt.Errorf("%d was provided: %w", answer, ErrIncorrectAnswer)
	}

	return nil

}

func main() {

	separator := "------------------------"

	// ASK USER A MATH QUESTION UNTIL THEY TYPE "stop"
	for {

		// ASK AND GET ANSWER
		userAnswer, err := askQuestion(os.Stdin)
		if err != nil {
			log.Fatalf("Error getting User Input: %s", err)
		}

		// STOP
		if userAnswer == "stop" {
			fmt.Println("Done")
			break
		}

		// CHECK USER ANSWER
		err = checkAnswer(userAnswer)
		if err != nil {
			switch {
			case errors.Is(err, strconv.ErrSyntax):
				fmt.Println("Not an integer")
			case errors.Is(err, ErrIncorrectAnswer):
				fmt.Println("INCORRECT!")
			}
			log.Errorf("Error with answer: %s\n%s\n", err, separator)
			continue
		}

		fmt.Printf("YOUR ANSWER %s IS CORRECT!\n%s\n", userAnswer, separator)

	}

}
