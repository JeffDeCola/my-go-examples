// my-go-examples errors.go

package main

import (
	"fmt"

	errors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func a(i int) (int, error) {

	// Do stuff

	// Call another function
	o, err := b(i)

	return o, err

}

func b(i int) (int, error) {

	// Do stuff

	// Call another function
	o, err := c(i)

	return o, err

}

func c(i int) (int, error) {

	// Do stuff
	fmt.Print("What is 2 + 2?")
	_, err := fmt.Scan(&answer)
	if err != nil {
		return "", fmt.Errorf("unable to open paraphrase: %w", err)
	}

	if i == 42 {
		// Make your error
		err := errors.New("can't work with 42")
		return 22, err
	}

	return i + 3, nil

}

func main() {

	//MULTIPLE RETURN VALUES
	r, err := a(43)
	if err != nil {
		// Handle the error
		log.Fatalf("Error getting paraphrase: %s", err)
		// log.Fatal("ERROR:", err)
	}
	// Continue
	fmt.Println("Returned", r)

	//MULTIPLE RETURN VALUES (Simpler form)
	r, err = a(42)
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatalf("Error getting paraphrase: %s", err)

		log.Fatal("ERROR:", err)
	}
	fmt.Println("Returned", r)
}
