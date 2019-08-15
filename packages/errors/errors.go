package main

import (
	"fmt"

	errors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func a(x int) (int, error) {
	if x == 42 {
		// Make your error
		return -1, errors.New("can't work with 42")
	}
	return x + 3, nil
}

func b(i int) (int, error) {

	// Do stuff

	// Call another function
	return a(i)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func main() {

	//MULTIPLE RETURN VALUES
	r, err := b(43)
	if err != nil {
		// Handle the error
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
	// Continue
	fmt.Println("Returned", r)

	//MULTIPLE RETURN VALUES (Simpler form)
	r, err = b(42)
	checkErr(err)
	fmt.Println("Returned", r)
}
