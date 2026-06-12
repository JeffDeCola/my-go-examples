// constructor-with-error
//
// Expanding on constructor-simple to add a configuration struct
// and error handling.

package main

import (
	"errors"
	"fmt"
	"os"
)

type config struct {
	a int
	s string
}

type thing struct {
	a        int
	s        string
	hostname string
}

func newThing(c config) (*thing, error) {

	// Validate
	if c.a < 0 {
		return nil, fmt.Errorf("a must not be negative, got %d", c.a)
	}
	if c.s == "" {
		return nil, errors.New("s must not be empty")
	}

	// Build
	a := c.a
	s := c.s
	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("hostname failed: %w", err)
	}

	// Assemble
	t := thing{
		a:        a,
		s:        s,
		hostname: hostname,
	}

	// Return
	return &t, nil

}

func printThing(t *thing) {
	fmt.Printf("a is %v, s is %q, hostname is %q\n", t.a, t.s, t.hostname)
}

func main() {

	// PART 1 - using a good config
	goodCfg := config{
		a: 4,
		s: "happy",
	}

	t, err := newThing(goodCfg)
	if err != nil {
		fmt.Println("newThing failed:", err)
		os.Exit(1)
	}

	printThing(t)

	t.a = 5
	t.s = "monkey"
	printThing(t)

	// PART 2 - using a bad config
	badCfg := config{
		a: -3,
		s: "smile",
	}

	t, err = newThing(badCfg)
	if err != nil {
		fmt.Println("newThing failed:", err)
		os.Exit(1)
	}

}
