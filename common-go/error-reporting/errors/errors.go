// my-go-examples errors.go

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var ErrIncorrectAnswer = errors.New("the answer is incorrect")

func a(answer string) error {

	// CALL ANOTHER FUNCTION - THIS IS ALL I DO
	err := b(answer)
	if err != nil {
		return fmt.Errorf("error calling b from a: %w", err)
	}

	return err

}

func b(answer string) error {

	// CONVERT TO INT
	answerInt, err := strconv.Atoi(answer)
	if err != nil {
		return fmt.Errorf("unable to convert in b: %w", err)
	}

	// CHECK ANSWER
	err = c(answerInt)
	if err != nil {
		return fmt.Errorf("error calling c from b: %w", err)
	}

	return err

}

func c(answer int) error {

	// IS ANSWER CORRECT?
	if answer != 4 {
		return fmt.Errorf("%d was provided: %w", answer, ErrIncorrectAnswer)
	}

	return nil

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	question := "What is 2+2?: "
	separator := "------------------------"

	// ASK USER A MATH QUESTION UNTIL THEY TYPE "stop"
	for fmt.Print(question); scanner.Scan(); fmt.Printf("%s\n%s", separator, question) {

		userAnswer := scanner.Text()

		if userAnswer == "stop" {
			fmt.Println("Done")
			break
		}

		err := a(userAnswer)
		if err != nil {
			switch {
			case errors.Is(err, strconv.ErrSyntax):
				fmt.Println("Not an int")
			case errors.Is(err, ErrIncorrectAnswer):
				fmt.Println("INCORRECT!!")
			}
			log.Errorf("Error with answer: %s", err)
			continue
		}

		fmt.Printf("YOUR ANSWER %s IS CORRECT!\n", userAnswer)

	}

}
